# GKE demo - Scaling apps using GKE autopilot

## Index
- [GKE demo - Scaling apps using GKE autopilot](#gke-demo---scaling-apps-using-gke-autopilot)
  - [Index](#index)
  - [Enable your GCP free trial](#enable-your-gcp-free-trial)
  - [Deploy a GKE cluster](#deploy-a-gke-cluster)
    - [Enable the required GCP services](#enable-the-required-gcp-services)
    - [Compute API](#compute-api)
    - [GKE API](#gke-api)
    - [Creation of a GKE Cluster](#creation-of-a-gke-cluster)
  - [Deploy the demo app](#deploy-the-demo-app)
    - [Check the IP allocated to the service](#check-the-ip-allocated-to-the-service)
  - [Simulate load on the app](#simulate-load-on-the-app)
    - [Install the load simulator tool](#install-the-load-simulator-tool)


This repository contains the material used during the hands-on session of the Food for thought session of the 31st March 2022 at [Deliverect](https://www.deliverect.com)

To be able to reproduce the session, you will need to have access to a GKE cluster. Luckily, Google provides you with a free trial access and 300$ of credits to experiment with Google Cloud Platform.

Find how to enable your free trial in the rest of this document.



## Enable your GCP free trial

1. Create yourself a Google account if you do not already have one.
2. Go to [https://console.cloud.google.com/](https://console.cloud.google.com/) and accept the free trial ![free_trial](img/free_trial.png)
3. Configure your payment information. Don't worry, they wont charge you anything.
4. You are all set, Enjoy !

## Deploy a GKE cluster
### Enable the required GCP services

Before we can deploy a GKE cluster, we need to enable the required GCP APIs.

- Compute Engine
- Google Kubernetes Engine

### Compute API
In the [GCP console](https://console.cloud.google.com/), go to [VPC Network page](https://console.cloud.google.com/networking/networks).

You will be redirected to the Compute Engine API enablement page. Click the `Enable` button and wait for the service to be activated.

Once activated you will be redirected to the VPC Network page and find that a default networki already exists.

### GKE API

Now go to the [Google Kubernetes Engine](https://console.cloud.google.com/kubernetes) page and enable this API too.

### Creation of a GKE Cluster

You are now ready to deploy a cluster. Click the `create` button and select `GKE Autopilot`.

Choose a name and a location for your cluster. Default value are fine if you do not know what to select.

Regarding networking, we will deploy a public cluster to make our live easy. But keep in mind that for a production enviroment, it is probably best to use a privater cluster.

Click the create button and wait for the cluster to be ready. It can take a few minutes.

## Deploy the demo app

This repository contains the kubernetes manifest in the [gke folder](./gke).

Here is a diagram of what you are going to deploy:
![app architecture](img/Micro%20service%20architecture.png)

The application is very simple and composed of a single microservice exposed by a service of type LoadBalancer. We want to keep it simple so we are not dealing with ingress.

To deploy the application execute the following command from a cloud shell:

```
kubectl apply -f gke/service.yaml
kubectl apply -f gke/deployment-autopilot.yaml
```

### Check the IP allocated to the service

To be able to reach the micro service we need to check what IP has been assigned to the service by GCP. 

Execute:
```
kubectl get service hello
```

And check the value of the `EXTERNAL-IP` column:
```
NAME    TYPE           CLUSTER-IP   EXTERNAL-IP    PORT(S)          AGE
hello   LoadBalancer   10.41.0.74   34.79.181.17   8080:31640/TCP   71s
```

Now you can try to access the microservice by sending an HTTP request go it. You should get a response with a nice hello world message.

```
curl http://34.79.181.17:8080/hello           
Hello World :-)
```

## Simulate load on the app

In order to simulate traffic going to our app. We are going to use a little tool that will send a LOT of request to it. Then we will watch the behavior of the deployment under load.

### Install the load simulator tool
The tool to simulate load is called `vegeta`. To install it you need to have Go installed and execute the following command:

```
go install github.com/tsenart/vegeta@latest
```


We can now start sending request to our app. By changing the value of the `-rate` flag we can choose how much request are sent by seconds.

```
echo "GET http://34.79.181.17:8080/load_test" | vegeta  attack -rate 500 -timeout 1s -keepalive=false | tee results.bin | vegeta report -every 1s
```

```
echo "GET http://34.79.181.17:8080/load_test" | vegeta  attack -rate 500 -timeout 1s -keepalive=false | tee results.bin | vegeta report -every 1s -type='hist[0,50ms,100ms,200ms]'
```

