# Cahier des charges — UpcycleConnect

**Projet annuel 2025-2026 — ESGI 2A**
**Appel d'offres du 5 février 2026**

---

## 1. Objet et contexte

UpcycleConnect (créée en 2021) est une entreprise dédiée à la réduction des déchets et à la valorisation
des matériaux via l'upcycling. Sa plateforme Web met en relation particuliers, artisans, professionnels et
salariés pour donner une seconde vie aux objets et matériaux inutilisés, avec un suivi complet des projets
et de leur impact environnemental.

En pleine croissance (France et Europe), l'entreprise souhaite moderniser l'ensemble de son système
d'information : gestion applicative, site Web, architecture systèmes et réseaux.

Ce cahier des charges présente les besoins du projet et les différentes parties à réaliser. Le projet est organisé en trois lots :

1. **Lot 1 — Gestion des solutions applicatives**
2. **Lot 2 — Infrastructure système, réseau et sécurité dédiée**
3. **Lot 3 — Modules complémentaires**

### 1.1 Sites de l'entreprise

- **174, rue La Fayette, Paris 10e** : siège, site principal de l'infrastructure.
- **Montreuil** : entrepôt de remise à neuf des matériaux.
- **Suisse** : premier partenaire hors espace Schengen (dépôt relais).
- **Datacenter** : hébergement des données critiques (RGPD).
- **Paris 11e et 13e** : bâtiments conférenciers (salles de réunion).
- Annexes prévues en province (16e, Bourg-la-Reine, Ivry, etc.).

### 1.2 Direction

| Rôle | Nom |
|------|-----|
| PDG | Sylvain Levy |
| DRH | Pierre Chabrier |
| Directeur Commercial | Ronnand Peuplus |
| Directrice Marketing | Laink Terracid |
| DSI | Norman Thavaud |
| Happyness Manager (Montreuil) | Antoine Maclair |
| Directeur régional (Suisse) | Frédéric Molas |

---

## 2. Mission 1 — Gestion des solutions applicatives

### 2.1 Espace Particuliers

- **Annonces** : dépôt d'une annonce de don ou vente d'objets (validation par le service administratif).
- **Conteneurs** : demande de dépôt d'un objet en box ; le service Check valide (don/vente) puis envoie
  un code d'ouverture ; un code-barres permet à l'artisan/professionnel de récupérer l'objet.
- **Espace Conseils** : accès aux conseils et bonnes pratiques.
- **Catalogue** : services, formations, événements (achats possibles en ligne).
- **Upcycling Score** : mesure de la quantité de déchets évités et ressources économisées.
- **Planning personnel** : cours, services, réservations.
- **Tutoriel de première connexion** : overlay bloquant à la première connexion.
- Notifications des activités UpcycleConnect.

### 2.2 Espace Professionnels et Artisans

- Gestion des contrats d'abonnement, publicité et **facturation**.
- Accès aux annonces avec **achats possibles**.
- Récupération des objets dans les conteneurs (lecture du code-barres).
- Gestion, suivi et mise en avant des **projets d'upcycling**.
- Notifications push (ex. OneSignal).

### 2.3 Espace Salariés (animateurs, formateurs)

- Création et animation de formations/ateliers (validation par un responsable).
- Gestion du planning.
- Gestion des conseils et news.
- Suivi et modération des forums.

### 2.4 Back-office d'administration générale

- Gestion complète des utilisateurs et acteurs.
- Gestion des contrats, documents et codes divers.
- **Gestion financière** (cf. modèle économique, annexe 1 du sujet).
- Validation des événements et des annonces.
- Gestion du catalogue des offres.
- Gestion des notifications (particuliers & professionnels).
- Gestion des conteneurs/box.

### 2.5 Contraintes techniques de la mission 1

- **Application Web** multi-espaces : couplage JavaScript/PHP/frameworks/**API Go**.
- **Paiement** par Stripe.
- Génération automatique de documents au **format PDF** (doublon conservé et accessible).
- **Notifications push** (ex. API OneSignal).
- **Multilingue** : ajout de langues possible sans passer par Google et sans modification du code.
- **API Go** pour l'intégralité des traitements ; services Cloud validés par les donneurs d'ordre.
- **Serveur Web personnel** : démonstration sur serveur déployé (réécriture d'URL, codes d'erreur).
  Un site en localhost ne sera pas corrigé.

---

## 3. Mission 2 — Infrastructure système, réseau et sécurité

### 3.1 Documents demandés

- Un **GANTT** permettant de suivre l'avancement des missions 1 et 2.
- Un **Trello** permettant de répartir les tâches entre les membres du groupe et de suivre leur avancement.
- Le **cahier des charges** complet (missions 1 et 2).
- La **documentation technique** des installations.
- Les **devis** en amont et dimensionnement (TVA 20 %) :
  - coût homme/heure travaillé,
  - infrastructure système,
  - infrastructure réseau.

### 3.2 Site de Paris (siège)

- **Active Directory / Domain Controller / DNS** (Windows Server 2019 ou 2022).
- **GPO de base** : fond d'écran, mappage des lecteurs réseau, interdiction de changer de mot de passe,
  interdiction d'installer des applications.
- **Baie de stockage** (Linux/Windows/TrueNAS) hébergeant toutes les données et bases (incl. mission 1).
- **DMZ** : serveur mail Windows (envoi/réception, notifications mission 1).
- Déploiement du site en **containers Docker (compose)**.
- Accès au site : **externe** (via DMZ), **interne** (LAN salariés), **VPN distant** (télétravail).

#### VLAN (LAN)

| VLAN | Usage |
|------|-------|
| Direction | Équipe de direction |
| Marketing | Équipe marketing |
| Commercial | Équipe commerciale |
| RH | Ressources humaines |
| Informatique | DSI / équipe technique |
| Hmanager | Happyness Manager |
| Regional | Directeur régional |
| DMZ | Support de la mission 1 |

- **Cluster HA de 2 firewalls OPNSense** (imposé, non substituable) entre LAN et DMZ.
- Connexion au réseau **RNSH** selon le schéma en annexe.
- **VPN externe** (OpenVPN ou WireGuard).
- Supervision : **GLPI** (ticketing + inventaire) et outil de **supervision** (Zabbix/Nagios/PRTG…).
- **Backup** complet et incrémental (Veeam, Acronis…) des données de l'infrastructure virtuelle et de la mission 1.

### 3.3 Site de Montreuil

- **VPN site-to-site** vers Paris uniquement.
- **VPN client-to-site** (OpenVPN/WireGuard).
- **Cluster HA de 2 firewalls OPNSense** (imposé).

### 3.4 Site Suisse (hors Schengen)

Sécurité renforcée (RGPD) réservée aux personnels certifiés « Spécialiste Bloc Infra ».

**GPO strictes** :
- fond d'écran bloqué ;
- installation de logiciels réservée aux administrateurs ;
- déploiement automatisé des agents GLPI/supervision ;
- blocage du CMD ;
- blocage des ports USB ;
- blocage de l'accès au pare-feu Windows ;
- mappage des lecteurs partagés pour le directeur régional.

**Équipements** :
- **Cluster HA de 2 firewalls PFSense** (imposé, non substituable) ;
- **1 VPN site-to-site** vers Paris.

### 3.5 Datacenter

- Accès depuis Paris et la Moldavie via **VPN OpenVPN site-to-site** avec certificats robustes.
- **Cluster de 2 firewalls** (imposé) : 1 PFSense + 1 OPNSense.

### 3.6 Routing Network Upcycle Connect (RNSH)

- **Routage inter-VLAN**.
- **VPN IPSec over GRE** au niveau des routeurs.
- **RIP ou OSPF**.
- Règles de filtrage des paquets sur les firewalls.

### 3.7 Box 11e et 13e

- **OpenVPN site-to-site** reliant les deux arrondissements.
- Wifi dédié : « Wifi_13 » (13e) et « Wifi_11 » (11e).
- **Serveur téléphonique 3CX** + softphones pour les salles de réunion.

### 3.8 Cadre de réalisation

- Virtualisation : **Eve-NG** ou **GNS3**.
- Location possible de serveurs hébergés chez un tiers (max ~40 € / groupe), sinon infra locale ou répartie.
- Mission 1 et modules Dev/Sciences (mission 3) déployés **en production sur l'infrastructure de la mission 2**.

---

## 4. Mission 3 — Modules complémentaires

Un module par étudiant selon le bloc électif, dans l'ordre des modules.

### 4.1 Bloc Développement

1. **M1** — Application mobile Android pour professionnels/artisans (3 fonctionnalités essentielles).
2. **M2** — Application mobile Android pour particuliers (2 fonctionnalités + Upcycling Score).
3. **M3** — Base **NoSQL** de gestion des documents + accès simplifié dans l'espace adhérent.

### 4.2 Bloc Sciences (IA, requêtes SQL, 3D)

1. **M1** — Reporting IA/Data Mining : description des acteurs (graphiques), dashboards du succès des
   prestations ; **modèle de ML de classification** prédictif des prestations demandées ; dataset de
   500 lignes minimum.
2. **M2** — Montée en charge à **15 000 utilisateurs** : tableau des performances de 4 requêtes
   problématiques max, puis solutions d'optimisation (tableau comparatif).
3. **M3** — Système intelligent et optimisé de ramassage des objets dans les conteneurs (entrepôts).

### 4.3 Bloc Infra (sécurité WiFi, admin Windows/Linux, services Cloud)

1. **M1** — Wifi sécurisé : VLAN **WiFi_GUEST** (adhérents) + **WIFI_LAN** (employés), synchronisation
   AD/RADIUS, normes en vigueur, démonstration de récupération d'un mot de passe Wifi + solution.
2. **M2** — Redondance/PRA : DHCP Linux, DNS Linux, backup quotidien à 23h, backup mensuel le 30,
   **RODC** secondaire de l'AD de Paris.
3. **M3** — Bastion sécurisé **Teleport** au datacenter (accès distant 24/7/365 aux données RGPD) +
   application de gestion des bases RGPD des adhérents ayant fréquenté le site Suisse.

---

## 5. Exigences transverses et organisationnelles

### 5.1 Organisation

- Chaque groupe réalise les **missions 1, 2 et 3**.
- Deux responsables sont désignés : un chef de projet réseau et un chef de projet développement.
- Chaque membre doit participer aux missions 1 et 2, avec une partie réseau et une partie développement.
- **Trello** obligatoire pour la gestion de projet.
- **GitHub** obligatoire, chaque membre du groupe doit avoir son **propre compte**.

### 5.2 Techniques

- Respect strict des langages de programmation imposés.
- Produit **packagé et déployable** : script d'installation/copie des répertoires, bibliothèques, fichiers
  et bases de données.
- Frameworks, API et bibliothèques autorisés, sous validation des responsables.
- Environnements de virtualisation autorisés : **Eve-NG**, **GNS3**.

### 5.3 Ce qui doit être rendu

| Livrable | Mission |
|----------|---------|
| Cahier des charges complet | 1 & 2 |
| GANTT | 1 & 2 |
| Trello (répartition équitable) | 1 & 2 |
| Documentation technique | 2 |
| Devis (homme/heure, système, réseau) avec TVA 20 % | 2 |
| Application Web (particuliers, pros, salariés) | 1 |
| Back-office d'administration | 1 |
| API Go + services Cloud | 1 |
| Infrastructure réseau/système/sécurité déployée | 2 |
| Module(s) complémentaire(s) Mission 3 | 3 |

---

## 6. Calendrier prévisionnel

*Les dates précises des différents rendus seront complétées en fonction des jalons indiqués sur MyGes.*

| Phase | Contenu |
|-------|---------|
| Initialisation | Cadrage, planification Trello/GANTT, cahier des charges |
| Conception | Architecture applicative (API Go, BDD, front) et réseau |
| Développement Mission 1 | Espaces particuliers, pros, salariés, back-office |
| Infrastructure Mission 2 | Sites Paris, Montreuil, Suisse, datacenter, RNUC, DMZ |
| Mission 3 | Modules par bloc électif |
| Intégration & production | Déploiement Docker sur Nginx, supervision, backups |
| Recette & livraison | Tests, documentation, packaging, démonstration |

---

## 7. Modèle économique (référence pour la gestion financière)

1. **Abonnements/Freemium** : 15–30 €/mois (particuliers gratuits ; pro/artisans premium : tableaux de
   bord avancés, analyse d'impact, statistiques matériaux, alertes collecte priorisées).
2. **Commissions** : 5–10 % sur les ventes d'objets/matériaux.
3. **Publicité & partenariats** : 100–500 €/mois/campagne (sponsoring de créateurs, marques éco-responsables).
4. **Formations & ateliers** : 20–100 €/personne (présentiel ou en ligne, abonnements pédagogiques).

> Remarque : les tarifs pourront évoluer d'une année à l'autre.