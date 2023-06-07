## Build and Run locally

Build:

```bash
docker build -t webapp:v1 . --platform=linux/amd64
```

Run locally:

```bash
docker run -e "ASPNETCORE_URLS=http://*:5002" -p 5002:5002 -it webapp:v1
```

Send a test request:

```bash
curl http://localhost:5002/api/invoketask?task=Task1
```

## Deploy to Azure Functions

Push image to ACR:

```bash
docker login <ACR_NAME>.azurecr.io
docker tag webapp:v1 <ACR_NAME>.azurecr.io/apps/webapp:v1
docker push <ACR_NAME>.azurecr.io/apps/webapp:v1
```

Deploy image to Azure Functions:

```bash
az functionapp config container set \
  --docker-custom-image-name --docker-custom-image-name <ACR_NAME>.azurecr.io/apps/webapp:v1 \
  --docker-registry-server-password <SECURE_PASSWORD> \
  --docker-registry-server-user <USER_NAME> \
  --name <APP_NAME> \
  --resource-group <RESOURCE_GROUP>
```

## Sample Workflow

```yaml
id: deploymentplan
name: Sample Deployment Plan
version: '1.0'
specVersion: '0.8'
states:
  - name: InvokeTask1
    type: operation
    actions:
      - name: Invoke InvokeTask1 Function
        functionRef: Task1
    transition: InvokeTask2
  - name: InvokeTask2
    type: operation
    actions:
      - name: Invoke Task2 Function
        functionRef: Task2
    transition: InvokeTask3
  - name: InvokeTask3
    type: operation
    actions:
      - name: Invoke Task3 Function
        functionRef: Task3
    end: true
functions:
  - name: Task1
    type: rest
  - name: Task2
    type: rest
  - name: Task3
    type: rest
```
