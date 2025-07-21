# 📲 Push Notifications Service

This microservice is responsible for sending push notifications across the platform, such as appointment confirmations, reminders, and system alerts. Built in **Go**, it follows a modular architecture with proper layering and supports monitoring via **Grafana** dashboards.

---

## 🧩 Features

- Send push notifications to users  
- REST API built with Go  
- Modular architecture: config, handlers, middleware, services  
- Integrated Grafana dashboards for observability  
- Docker-ready  
- CI/CD with GitHub Actions  
- Deployable to AWS EC2

---

## 📁 Key Folders

- `cmd/main.go`: Entry point of the application  
- `internal/`: Application logic
  - `config/`: Loads env variables
  - `handlers/`: HTTP request logic
  - `middleware/`: Auth, logging, etc.
  - `services/`: Push logic (e.g., FCM, SNS)
- `grafana/dashboards/push.json`: Monitoring dashboard for Grafana

---

## ⚙️ Environment Variables (`.env`)

```env
PORT=4005
PUSH_API_KEY=your_push_api_key
PUSH_SERVICE=FCM
GRAFANA_ENABLED=true

    You can extend support to more push services by adding to services/.

🐳 Docker Usage
1. Build the image

docker build -t push-notifications .

2. Run the container

docker run -d -p 4005:4005 --env-file .env push-notifications

🔌 API Endpoints
Method	Route	Description
POST	/push/send	Sends a push notification

    Example JSON payload:

{
  "title": "Appointment Reminder",
  "message": "You have an appointment tomorrow at 10 AM",
  "recipient": "user123"
}

☁️ EC2 Deployment
1. SSH into your EC2 instance

ssh -i key.pem ec2-user@<YOUR_EC2_PUBLIC_IP>

2. Pull the image

docker pull jeffri1997/push-notifications:qa

3. Run the service

docker run -d -p 4005:4005 --env-file .env jeffri1997/push-notifications:qa

✅ Make sure port 4005 is open in your EC2 security group.
📊 Grafana Dashboard

The included grafana/dashboards/push.json file can be imported into Grafana to visualize push activity.
How to use:

    Open Grafana

    Go to Dashboards > Import

    Upload push.json

    Assign the correct data source (e.g., Prometheus)

🤖 GitHub Actions CI/CD

This workflow automates build and push to Docker Hub on every push to qa:

name: Build and Push Push Notifications Service

on:
  push:
    branches: [qa]

jobs:
  build-and-push:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - name: Docker Login
        uses: docker/login-action@v2
        with:
          username: ${{ secrets.DOCKER_USERNAME }}
          password: ${{ secrets.DOCKER_PASSWORD }}

      - name: Build and Push
        run: |
          docker build -t jeffri1997/push-notifications:qa .
          docker push jeffri1997/push-notifications:qa

🧪 Testing

You can run tests (if added) using:

go test ./...

👤 Author

Jefferson Marcalla
GitHub: @Jeff97ares