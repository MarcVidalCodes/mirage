# Mirage 
###*IN DEVELOPMENT*

Mirage is a Kubernetes operator that automatically provisions isolated full-stack preview environments for every GitHub PR. It treats infrastructure as a disposable resource, eliminating the "it works on my machine" issue. 

## Features / Roadmap
- Dynamic Provisioning: Automatically spins up K8s Namespaces and Deployments based on PR content.
- Ingress Automation: Generates dynamic URLs (e.g., pr-101.mirage.internal) for immediate preview.
- TTL Reaper: Automatically destroys environments after X hours of inactivity to save cloud costs.
- Drift Detection: Watches for manual changes and reverts them to ensure GitOps compliance.

## Why I built it / The Vision
I building Mirage to solve the "Environment Drift" that plagued my previous internships, where code that worked locally or in staging worked differently than it did in production. By spinning up production like environments for every Pull Request, Mirage shifts infrastructure testing, ensuring that if a feature works in teh PR, it works in the real world. It tranforms infrastructure from a static bottleneck into a disposable resource that emppowers devs to test without fear. 

## Getting Started
1. ```git clone https://github.com/MarcVidalCodes/mirage.git```

2. ```cd mirage-operator```

3. Run the server 

    ```go run main.go```

4. Simulate a WebHook
    ```curl -X POST http://localhost:8080/webhook \
     -H "Content-Type: application/json" \
     -d '{ "action": "opened", "number": 101, ... }'```


## Tech Stack
- Golang (go1.25.6)
- Kubernetes
- Docker
- AWS EKS (Planned)