import * as pulumi from '@pulumi/pulumi'
import { Config, Cluster, Dockerhub } from '@shapeshiftoss/common-pulumi'

const getStorageClass = (cluster: string) => {
  switch (cluster) {
    case 'docker-desktop':
      return 'hostpath'
    case 'minikube':
      return 'standard'
    case 'eks':
      return 'gp2'
    default:
      throw new Error(`cluster not supported... Use: 'docker-desktop', 'minikube', or 'eks.`)
  }
}

export interface EthereumConfig {
  kubeconfig: string
  config: Config
  namespace: string
}

export const getConfig = async (): Promise<EthereumConfig> => {
  const config = new pulumi.Config('unchained').requireObject<Config>('ethereum')
  const stackReference = new pulumi.StackReference(config.stack)
  const kubeconfig = (await stackReference.getOutputValue('kubeconfig')) as string
  const namespaces = (await stackReference.getOutputValue('namespaces')) as Array<string>
  const defaultNamespace = (await stackReference.getOutputValue('defaultNamespace')) as string

  const namespace = config.environment ? `${defaultNamespace}-${config.environment}` : defaultNamespace
  if (!namespaces.includes(namespace)) {
    throw new Error(
      `Error: environment: ${config.environment} not found in cluster. Either remove to use default environment or verify environment exists`
    )
  }

  config.isLocal = (await stackReference.getOutputValue('isLocal')) as boolean
  config.cluster = (await stackReference.getOutputValue('cluster')) as Cluster
  config.dockerhub = (await stackReference.getOutputValue('dockerhub')) as Dockerhub
  config.rootDomainName = (await stackReference.getOutputValue('rootDomainName')) as string

  const missingRequiredConfig: Array<string> = []

  if (!config.stack) missingRequiredConfig.push('stack')

  if (config.mongo) {
    config.mongo.storageClass = getStorageClass(config.cluster)

    if (!config.mongo.cpuLimit) missingRequiredConfig.push('mongo.cpuLimit')
    if (!config.mongo.helmChartVersion) missingRequiredConfig.push('mongo.helmChartVersion')
    if (!config.mongo.memoryLimit) missingRequiredConfig.push('mongo.memoryLimit')
    if (!config.mongo.replicaCount) missingRequiredConfig.push('mongo.replicaCount')
    if (!config.mongo.storageSize) missingRequiredConfig.push('mongo.storageSize')
  }

  if (config.indexer) {
    config.indexer.storageClass = getStorageClass(config.cluster)

    if (!config.indexer.cpuLimit) missingRequiredConfig.push('indexer.cpuLimit')
    if (!config.indexer.memoryLimit) missingRequiredConfig.push('indexer.memoryLimit')
    if (!config.indexer.replicas) missingRequiredConfig.push('indexer.replicas')
    if (!config.indexer.storageSize) missingRequiredConfig.push('indexer.storageSize')
  }

  if (config.indexer?.daemon) {
    config.indexer.daemon.storageClass = getStorageClass(config.cluster)

    if (!config.indexer.daemon.cpuLimit) missingRequiredConfig.push('indexer.daemon.cpuLimit')
    if (!config.indexer.daemon.image) missingRequiredConfig.push('indexer.daemon.image')
    if (!config.indexer.daemon.memoryLimit) missingRequiredConfig.push('indexer.daemon.memoryLimit')
    if (!config.indexer.daemon.storageSize) missingRequiredConfig.push('indexer.daemon.storageSize')
  }

  if (missingRequiredConfig.length) {
    throw new Error(
      `Missing the following configuration values from Pulumi.${pulumi.getStack()}.yaml: ${missingRequiredConfig.join(
        ', '
      )}`
    )
  }

  return { kubeconfig, config, namespace }
}