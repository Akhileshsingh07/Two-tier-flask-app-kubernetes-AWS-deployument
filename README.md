Two-Tier Flask & MySQL Application Deployment on Kubernetes (AWS EKS) with Docker Compose


This repository contains the source code, Docker Compose configuration, and Kubernetes manifests for deploying a two-tier application on AWS EKS. The application consists of a Flask web server and a MySQL database, both containerized with Docker and orchestrated using Kubernetes. This README provides an overview of the project and instructions for deployment.


Overview

This project demonstrates the deployment of a two-tier application, with the following components:

    Flask Web Server: A Python-based web server using the Flask framework. This server serves a simple web application that interacts with the MySQL database.

    MySQL Database: A containerized MySQL database used as the backend data store for the web application.

    Kubernetes (AWS EKS): Kubernetes is used to orchestrate the deployment of the application components. AWS EKS (Elastic Kubernetes Service) is the chosen platform for Kubernetes cluster management.

    Docker Compose: Docker Compose is used for local development and testing of the application components.


Project Structure

The project's directory structure is organized as follows:

/
├── app/
│   ├── Dockerfile
│   ├── ...
├── db/
│   ├── Dockerfile
│   ├── ...
├── k8s/
│   ├── deployment.yaml
│   ├── service.yaml
│   ├── ...
├── docker-compose.yml
├── README.md

    app/: Contains the Flask application code and a Dockerfile for building the Flask web server image.

    db/: Contains the MySQL database initialization script and a Dockerfile for building the MySQL image.

    k8s/: Contains Kubernetes manifest files for deploying the application to an AWS EKS cluster.

    docker-compose.yml: Defines Docker Compose services for local development and testing.

Configuration

To configure and customize the application and deployment for your specific needs, follow these steps:

    Clone the repository to your local machine.

    Modify the Flask application code in the app/ directory as needed.

    Update the MySQL initialization script in the db/ directory if necessary.

    Customize the Kubernetes manifests in the k8s/ directory to match your deployment requirements.

    Set the environment-specific configuration in the Kubernetes manifests, such as database credentials, service names, and resources.

Deployment

To deploy the application on AWS EKS with Kubernetes and Docker Compose, follow these steps:

    Build the Docker images for the Flask application and MySQL database:

    shell

docker-compose build

Deploy the application locally for testing using Docker Compose:

shell

docker-compose up

Deploy the application to AWS EKS:

    Apply the Kubernetes manifests in the k8s/ directory:

    shell

    kubectl apply -f k8s/

    Monitor the deployment using kubectl.

Access the application by obtaining the external IP address or hostname provided by the Kubernetes service.
