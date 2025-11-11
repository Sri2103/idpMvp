# 🧩 Internal Developer Platform (IDP) — MVP Monorepo

> A self-service platform enabling developers to provision, deploy, and monitor applications with minimal friction — built with a modular microservice architecture and decentralized deployments.

***

## 🚀 Overview

The **IDP MVP** provides a foundation for managing tenants, syncing with GitHub and ArgoCD, generating kubeconfigs, and visualizing service health — while allowing seamless scaling into a production-grade system later.

This repository is structured as a **monorepo** to share core packages and streamline local development, while still allowing **independent deployments per service**.

***

## 🏗️ Core Services

| Service | Description | Technology |
| :-- | :-- | :-- |
| **API Gateway** | Entry point for all clients (UI, CLI). Handles routing, authentication, and rate limiting. | Go / Gin |
| **Auth Service** | JWT \& OIDC authentication, tenant onboarding, token issuance, and service account management. | Go / client-go / OIDC |
| **Tenant Service** | Manages tenants, namespaces, quotas, and access controls within the platform. | Go / Kubernetes API |
| **Repo Service** | Integrates with GitHub / GitLab to manage repositories, webhooks, and commit metadata. | Go / GitHub API |
| **ArgoCD Sync Service** | Syncs tenant applications with ArgoCD projects and tracks deployment states. | Go / ArgoCD API |
| **Kubeconfig Service** | Dynamically generates per-tenant kubeconfigs for scoped cluster access. | Go / Kubernetes client-go |
| **Health Service** | Collects and serves runtime health, metrics, and event logs for UI and CLI visualization. | Go / Prometheus client |


***

## 🧱 Shared Packages (`/pkg`)

| Package | Purpose |
| :-- | :-- |
| `api` | Shared OpenAPI types, REST definitions, and generated Swagger specs. |
| `db` | PostgreSQL connection pools, migrations, and ORM helpers. |
| `config` | Environment variable and YAML configuration loaders. |
| `logger` | Centralized structured logging using Zerolog or Zap. |
| `k8s` | Kubernetes client helpers, CRD utilities, informer wrappers. |
| `middleware` | Common HTTP middlewares: CORS, JWT, tracing, panic recovery. |
| `utils` | Small utilities like string, file, and JSON helpers. |
| `events` | EventBus abstraction (NATS / Kafka for later production use). |
| `constants` | Shared constants and enums for inter-service consistency. |


***

## 🗂️ Repository Layout

```plaintext
idp/
├── cmd/
│   ├── api-gateway/
│   ├── auth-service/
│   ├── tenant-service/
│   ├── repo-service/
│   ├── argocd-service/
│   ├── kubeconfig-service/
│   └── health-service/
│
├── internal/
│   ├── auth/
│   ├── tenant/
│   ├── repo/
│   ├── argocd/
│   ├── kubeconfig/
│   └── health/
│
├── pkg/
│   ├── api/
│   ├── config/
│   ├── db/
│   ├── k8s/
│   ├── logger/
│   ├── middleware/
│   ├── utils/
│   └── events/
│
├── deployments/
│   ├── charts/
│   ├── manifests/
│   ├── Dockerfiles/
│   └── values/
│
├── infra/
│   ├── postgres/
│   ├── argocd/
│   └── k8s/
│
├── scripts/
│   ├── build.sh
│   ├── run_local.sh
│   └── test.sh
│
├── .github/workflows/
│   ├── build.yml
│   ├── test.yml
│   └── deploy.yml
│
├── .gitignore
├── .dockerignore
├── .helmignore
├── go.mod
├── go.sum
├── Makefile
└── README.md
```


***

## 🧩 Tech Stack

- **Language:** Go 1.23+
- **API:** REST + Swagger (OpenAPI 3.0)
- **Database:** PostgreSQL
- **Orchestration:** Kubernetes + ArgoCD
- **CI/CD:** GitHub Actions
- **Auth:** JWT / OIDC
- **Logging:** Zerolog / Zap
- **Metrics:** Prometheus + Grafana (optional)
- **Containerization:** Docker + Helm

***

## 🧰 Local Development

### Prerequisites

- Go ≥ 1.23
- Docker \& Docker Compose
- Minikube / Kind cluster (optional)
- kubectl \& helm
- PostgreSQL (local or containerized)

***

### Run locally

```bash
# Start dependencies
docker compose up -d postgres

# Run a service (e.g., auth-service)
make run AUTH
```


***

### API Documentation

Each service exposes `/swagger` endpoint.

Generated specs are stored under:

`/pkg/api/generated/swagger/`

To regenerate:

```bash
make gen-swagger
```


***

### Deployment

Each service is deployable independently using Helm charts:

```bash
helm install idp-auth deployments/charts/auth-service
```

All services share a base configuration under `/deployments/values/`.

CI/CD pipelines handle:

- Lint \& test per service
- Build \& push Docker image
- Deploy via ArgoCD manifests

***

### Security \& Secrets

Secrets are managed via Kubernetes Secrets (dev) or Vault (production).

`.env`, `.env.*`, and `values-secret.yaml` are excluded via `.gitignore`.

***

### Roadmap

| Stage | Description | Status |
| :-- | :-- | :-- |
| **MVP** | Core services + GitHub + ArgoCD integration | ✅ In progress |
| **CLI Client** | Command-line automation interface | 🔜 Next |
| **Observability** | Add metrics, events, and logs service | ⏳ Planned |
| **Policy Engine** | RBAC mapper, compliance hooks | ⏳ Planned |
| **Production** | Full horizontal scalability, caching, and async events | 🕓 Later |

