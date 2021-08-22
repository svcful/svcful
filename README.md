# Svcful: A tool for defining and maintaining a service-oriented architecture

## How it works

1. Define service specifications
    ```sh
    sf init .
    
    # follow cli instructions and a svcful.yaml should be created
    ```
2. Generate documentation from service specifications
    ```sh
    # generate documentation for the service named 'app'
    sf create docs --name app

    # generate documentation for all services
    sf create docs --all
    ```
3. Initialise repositories from service specifications (integrate with your source-code hosting platfrom to enable managing of children repositories)
    ```sh
    # optional step to connect to github, follow cli instructions
    sf connect github

    # optional step to connect to gitlab, follow cli instructions
    sf connect gitlab

    # generate directory for the service named 'app'
    sf create service --name app

    # generate directories for all known services
    sf create service --all

    # generate directories and repositories for all known services
    sf create service --all --repos
    ```
4. Generate contract tests between dependent services
5. Generate Terraform scripts to bring up required infrastructure
    ```sh
    sf create terraform
    ```
6. Generate Helm Charts to deploy services onto Kubernetes clusters
    ```sh
    sf create helm
    ```
7. Generate Ansible playbooks to deploy services onto standalone machines
    ```sh
    sf create ansible
    ```
8. Generate Docker Compose manifests for local development requirements
    ```sh
    sf create docker-compose
    ```
9.  Run model-based testing to identify system failure patterns
    ```sh
    sf test
    ```

## Example service specification

```yaml
name: my-ecommerce-shop
services:
  - name: app
    description: this service is the mobile app
    services:
      - ref: cart
        description: for updating cart
      - ref: checkout
        description: for handling user checkouts
      - ref: newsfeed
        description: for displaying news to user
      - ref: payments
        description: for processing user payment
      - ref: security
        description: for authenticating users
      - ref: store
        description: for showing users available products
  - name: cart
    description: this service handles customer shopping carts
    services:
      - type: database
        description: for storing cart data
        database:
          name: mongo
  - name: checkout
    description: this service hanldes customer checkouts
    services:
      - ref: cart
        description: for retreiving items in cart
      - ref: security
        description: for retrieving user account info
      - ref: store
        description: for checking if items in cart are still available
  - name: inventory
    description: this is an external service provider that hanldes the inventory
    type: provider
  - name: newsfeed
    description: this service handles the news
    services:
      - type: database
        description: for storing blogs/articles
        database:
          name: mongo
  - name: news-updater
    description: this service updates the newsfeed
    lifecycle:
      type: cronjob
  - name: notifications
    description: this service handles user notifications
    services:
      - ref: security
        description: for retrieving user account details
  - name: payment-provider
    description: this service interfaces with the financial institute
    type: provider
  - name: payments
    description: this service hanldes finanncial transactions
    services:
      - ref: security
        description: for verifying user account information
      - ref: payment-provider
        description: for handling the actual payment processing
      - type: database
        database:
          name: postgres
  - name: security
    description: this services handles user authentication
    services:
      - ref: notifications
        description: for sending notifications if security-related incidents are found
      - type: database
        database:
          name: postgres
  - name: store
    description: this service handles products on sale
    services:
      - ref: inventory
      - type: database
        database:
          name: postgres
  - name: website
    description: this service is the web-accessible frontend
    type: public
    services:
      - ref: cart
      - ref: checkout
      - ref: newsfeed
      - ref: payments
      - ref: security
      - ref: store
```

## Updating service specifications

1. For services that are part of the repository, configuration will be directly updated. Merge the updated specifications to complete the specification change
2. For services that are in their own repositories, a branch will be created storing the configuration changes. Merge the updated specifications in the service's repository to complete the specification change
