param(
    [Parameter(Mandatory=$true)]
    [String]
    $AKS_CLUSTER_RESOURCE_GROUP,
    [Parameter(Mandatory=$true)]
    [String]
    $AKS_CLUSTER_NAME,
    [Parameter(Mandatory=$true)]
    [String]
    $AZURE_CONTAINER_REGISTRY_RESOURCE_GROUP,
    [Parameter(Mandatory=$true)]
    [String]
    $AZURE_CONTAINER_REGISTRY_NAME
)

$NODE_RESOURCE_GROUP = az aks show `
	--resource-group $AKS_CLUSTER_RESOURCE_GROUP `
	--name $AKS_CLUSTER_NAME `
	--query "nodeResourceGroup" `
	-o tsv

$AGENT_POOL_NAME = "$AKS_CLUSTER_NAME-agentpool"

$AGENT_POOL_IDENTITY = az identity show `
	--name $AGENT_POOL_NAME `
	--resource-group $NODE_RESOURCE_GROUP `
	--query "principalId" `
	-o tsv

$ACR_ID = az acr show `
	--name $AZURE_CONTAINER_REGISTRY_NAME `
	--resource-group $AZURE_CONTAINER_REGISTRY_RESOURCE_GROUP `
	--query "id" `
	-o tsv

az role assignment create `
	--role "AcrPull" `
	--assignee-principal-type "ServicePrincipal" `
	--assignee-object-id "$AGENT_POOL_IDENTITY" `
	--scope $ACR_ID