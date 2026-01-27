# AegisFlux  
**Kubernetes Network Monitor + Honeypot + API + Dashboard**

AegisFlux es una plataforma completa para observabilidad y seguridad en entornos Kubernetes.  
Incluye:

- ✔ Agente DaemonSet escrito en **Go**  
- ✔ Métricas de red (Prometheus exporter)  
- ✔ Honeypots HTTP y SSH  
- ✔ API central para gestionar eventos  
- ✔ Dashboard frontend (Angular)  
- ✔ Despliegue automatizado en Kubernetes  
- ✔ CI/CD mediante GitHub Actions o Jenkins  

---

## Arquitectura General

```mermaid
flowchart TD

    subgraph Cluster [Kubernetes Cluster]
        A1[DaemonSet - AegisFlux Agent<br>Go Exporter /metrics] --> P[Prometheus]

        H1[Honeypot HTTP<br>:8088] --> API
        H2[Honeypot SSH<br>:2222] --> API

        API[API Go /events] --> DB[(Memoria / Futuro DB)]

    end

    P --> D[Dashboard Angular]
    API --> D

