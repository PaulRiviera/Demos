param(
    [Parameter(Mandatory=$true)]
    [String]
    $REGION_NAME = "eastus2",
    [Parameter(Mandatory=$true)]
    [String]
    $RESOURCE_GROUP_NAME,
    [Parameter(Mandatory=$true)]
    [String]
    $LOG_ANALYTICS_NAME,
    [Parameter(Mandatory=$true)]
    [String]
    $AKS_CLUSTER_NAME,
    [String]
    $AKS_NETWORK_PLUGIN = "azure",
    [String]
    $AKS_NODE_POOL = "system",
    [String]
    $AKS_VM_SIZE = "Standard_B2s",
    [String]
    $NODE_COUNT = "2",
    [String]
    $VERSION
)

if ($VERSION -eq "") {
    $VERSION = az aks get-versions `
        --location $REGION_NAME `
        --query 'orchestrators[?!isPreview] | [-1].orchestratorVersion' `
        --output tsv
}

az group create `
    --name $RESOURCE_GROUP_NAME `
    --location $REGION_NAME

az monitor log-analytics workspace create `
    --resource-group $RESOURCE_GROUP_NAME `
    --workspace-name $LOG_ANALYTICS_NAME `
    --location $REGION_NAME

$LOG_ANALYTICS_ID = az monitor log-analytics workspace show `
    --resource-group $RESOURCE_GROUP_NAME `
    --workspace-name $LOG_ANALYTICS_NAME `
    --query "id" `
    -o tsv

[string[]]$CLUSTERS = az aks list `
    -g $RESOURCE_GROUP_NAME `
    --query "[].name" `
    -o tsv

if ($null -ne $CLUSTERS -and $CLUSTERS.Contains($AKS_CLUSTER_NAME)) {
    Write-Output "Updating Cluster"
    az aks update `
        --resource-group $RESOURCE_GROUP_NAME `
        --name $AKS_CLUSTER_NAME `
        --max-count 2 `
        --min-count 1 `
        --enable-cluster-autoscaler `
        --enable-aad
} else {
    Write-Output "Creating Cluster"
    az aks create `
        --resource-group $RESOURCE_GROUP_NAME `
        --name $AKS_CLUSTER_NAME `
        --workspace-resource-id $LOG_ANALYTICS_ID `
        --kubernetes-version $VERSION `
        --network-plugin $AKS_NETWORK_PLUGIN `
        --vm-set-type VirtualMachineScaleSets `
        --load-balancer-sku standard `
        --nodepool-name $AKS_NODE_POOL `
        --node-vm-size $AKS_VM_SIZE `
        --node-count $NODE_COUNT `
        --max-count 2 `
        --min-count 1 `
        --enable-addons monitoring `
        --enable-cluster-autoscaler `
        --enable-managed-identity `
        --generate-ssh-keys
}