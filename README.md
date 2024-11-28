# devops-workshop-p2

Lær hvordan du [deployer](https://teknisk-ordbok.fly.dev/ordbok/Deploy) koden din til [prod](https://teknisk-ordbok.fly.dev/ordbok/Produksjon)!

# Del 2: Docker og Kubernetes

## ▶️ 3. Docker

### 🔨 Oppgave 3.1

Prøv å bygg frontend med Docker lokalt med denne kommandoen:

```bash
docker build frontend -t frontend:latest
```

Den kommer til å feile. Les feilmeldingen og oppdater [base image](https://docs.docker.com/build/building/base-images) til et som er kompatibelt.

💡 _HINT:_ Se på dokumentasjonen til [Node.js sine images](https://hub.docker.com/_/node).

<details>
  <summary>✨ Se fasit</summary>

```dockerfile
FROM node:22-alpine # endre til nyere versjon av Node.js

WORKDIR /app

COPY yarn.lock package.json ./
RUN yarn install --frozen-lockfile

COPY index.html ./

CMD ["yarn", "start"]
```

</details>

### 🔨 Oppgave 3.2

Kjør Docker-containeren du nettop bygde med denne kommandoen:

```bash
docker run -it frontend:latest -p 3000:3000
```

Da skal du kunne gå i nettleseren på denne linken: http://localhost:3000

**Ignorer at den ikke får kontakt med backend, det fikser vi i neste oppgave**

### 🔨 Oppgave 3.3

Vi skal bygge backend'en vår med Docker.
Når vi kjører flere tjenester er det lettere å bruke **Compose**.
Det gjør at vi kan koble sammen flere containere og kjøre de samtidig.

Bruk denne kommandoen for å kjøre frontend og backend med Docker Compose:

```bash
docker compose up --build
```

Du kan gå til http://localhost:3000 for å se frontend og http://localhost:8080/api for å se backend.
Du vil også se at frontend fortsatt ikke har kontakt med backend.

Vi definerer hvilken port frontend-containeren skal åpne opp i [docker-compose.yml](docker-compose.yml).
Gjør dette for backend også. Når du har gjort det, kan du kjøre `docker compose up --build` igjen.

💡 _HINT:_ Se på porten som backend bruker i [main.go](backend/main.go).

<details>
  <summary>✨ Se fasit</summary>

```yaml
services:
  frontend:
    build: frontend
    ports:
      - "3000:3000"
    links:
      - backend

  backend:
    build: backend
    # legg til riktig port
    ports:
      - "8080:8080"
```

</details>

### 🔨 Oppgave 3.4

Vi kan definere miljøvariabler i Docker og Docker Compose.
Vi har en miljøvariabel i [main.go](backend/main.go) som heter `MESSAGE`.
Ved å overskrive denne kan vi endre hva som blir vist i frontend.

Prøv å sette `MESSAGE`-variabelen i [docker-compose.yml](docker-compose.yml) og kjør `docker compose up --build` igjen.

💡 _HINT:_ Se på dokumentasjonen til [Compose](https://docs.docker.com/compose/how-tos/environment-variables/set-environment-variables).

<details>
  <summary>✨ Se fasit</summary>

```yaml
services:
  frontend:
    build: frontend
    ports:
      - "3000:3000"
    links:
      - backend

  backend:
    build: backend
    ports:
      - "8080:8080"
    # legg til miljøvariabel
    environment:
      MESSAGE: "Min kule melding"
```

</details>

## 🏗️ 4. Kubernetes

### 📖 Før du begynner

Sjekk ut en git branch med navnet ditt, f.eks.:

```bash
git checkout -b andreas-bakseter
```

**PASS PÅ AT INGEN ANDRE HAR EN BRANCH MED SAMME NAVN SOM DEG!**

Hver gang du vil teste endringer, push de til branchen din:

```bash
git push -u origin andreas-bakseter # første gang
git push # senere
```

...og lag en pull request mot `master`-branchen.

Da vil GitHub Actions (prøve å) bygge Docker image og deploye til Kubernetes for deg.

### 🔨 Oppgave 4.1

Vi har lyst til å deploye vår backend til Kubernetes.
Definisjonen for Namespace, Deployment og Service er allerede laget for deg i [resources.yml](backend/resources.yml).

Åpne en pull request mot `master`-branchen og se om GitHub Actions bygger og deployer for deg.
Du burde få en IP-addresse m/ portnummber til backend'en din som output i GitHub Actions.

Gå til addressen i nettleseren din og se om du får svar fra backend.
Full addresse vil være `http://<IP med port>/api`.

### 🔨 Oppgave 4.2

Noe av fordelen med Kubernetes er at vi kan skalere opp og ned tjenestene våre veldig enkelt.
Prøv å endre antall replicas i [resources.yml](backend/resources.yml) fra 1 til 2 eller 3 (**NB: IKKE ØK ANTALL REPLICAS OVER 3**, takk).

Når du går inn på IP-adressen din igjen, vil du se at du får svar fra to forskjellige pods.
*Det kan hende nettleseren din cacher eller gjenbruker samme forbindelse, så prøv å åpne en inkognitotab eller terminalen:*

```bash
curl -v http://<IP med port>/api
```

Du kan også prøve å endre `backendHost` og `backendPort` i [frontend/index.html](frontend/index.html) til å peke på IP-adressen til Kubernetes.

### 🔨 Oppgave 4.3

Skaler ned antall replicas til 1 igjen.

Nå skal vi overskrive miljøvariabelen `MESSAGE` ved å sette den i [resources.yml](backend/resources.yml).

Push endringen til branchen din, og se om backend'en din har fått ny melding.

<details>
  <summary>✨ Se fasit</summary>

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: <DEPLOYMENT_NAME>
  name: <DEPLOYMENT_NAME>
  namespace: <NAMESPACE_NAME>
spec:
  replicas: 1
  selector:
    matchLabels:
      app: <DEPLOYMENT_NAME>
  strategy: {}
  template:
    metadata:
      labels:
        app: <DEPLOYMENT_NAME>
    spec:
      containers:
      - image: <IMAGE_NAME>
        name: <DEPLOYMENT_NAME>
        ports:
        - containerPort: 8080
        env:
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        # legg til miljøvariabel
        - name: MESSAGE
          value: "Min kule melding fra Kubernetes!"
        resources:
          limits:
            cpu: 30m
            memory: 32Mi
          requests:
            cpu: 30m
            memory: 32Mi
```

</details>

## 🏁 Ferdig!

Lukk pull requesten din.
Da vil GitHub Actions rydde opp etter seg og slette ressursene som ble opprettet i Kubernetes.

# 🤓 Setup for spesielt interesserte (ikke en del av workshop'en)

1. Lag en profil på [Hetzner Cloud](https://www.hetzner.com/cloud).
2. Sett opp et k3s-cluster på Hetzner Cloud ved hjelp av [Kube-Hetzner](https://github.com/kube-hetzner/terraform-hcloud-kube-hetzner). Følg README'en i repoet.
3. Få ut kubeconfig-output fra Terraform og paste den inn i GitHub Secrets med navnet `KUBECONFIG`.
4. Gjør noe liknende som i [deploy.yml](.github/workflows/deploy.yml) for å deploye til ditt eget cluster.
