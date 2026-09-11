# DEVIS — Modernisation du système d'information UpcycleConnect

| | |
|---|---|
| **Prestataire** | AlterNext |
| **Client** | UpcycleConnect — 174, rue La Fayette, 75010 Paris |
| **Appel d'offres** | 5 février 2026 — Projet annuel ESGI 2A 2025-2026 |
| **Objet** | Refonte complète du SI : solutions applicatives (M1), infrastructure système/réseau/sécurité (M2), modules complémentaires (M3) |
| **Date du devis** | 11 septembre 2026 |
| **TVA** | 20 % (France) |

> Devis établi conformément au cahier des charges (documents demandés, § 3.1 et § 5.3) :
> coût homme/heure travaillé, infrastructure système, infrastructure réseau.

---

## Récapitulatif général

| Devis | Libellé | Total HT |
|---|---|---|
| Devis n°1 | Coût des ressources humaines (homme/jour — TJM) | 69 525,00 € |
| Devis n°2 | Infrastructure système | 34 200,00 € |
| Devis n°3 | Infrastructure réseau | 29 000,00 € |
| **Total HT** | | **132 725,00 €** |
| TVA 20 % | | **26 545,00 €** |
| **Total TTC** | | **159 270,00 €** |

---

# DEVIS N°1 — COÛT RESSOURCES HUMAINES (homme/jour, TJM)

**Estimation globale : 93,5 jours-homme** (soit ~748 heures travaillées), détaillés ci-dessous.

> Le **TJM** (tarif journalier moyen) est facturé à la journée, sur la base d'une journée de travail de 8 h.
> TJM moyen pondéré estimé du projet : **≈ 744 € HT / jour**.

## Charge par mission (jours-homme)

| Mission | Contenu | Jours-homme |
|---|---|---|
| Pilotage | Cadrage, cahier des charges, GANTT, Trello, suivi, démonstrations | 7,5 |
| **Mission 1** | Application Web complète (particuliers, pros, salariés, back-office, API Go, Stripe, PDF, push, i18n) | 46,0 |
| **Mission 2** | Infrastructure système, réseau, sécurité (Paris, Montreuil, Suisse, datacenter, RNSH, box) | 32,5 |
| **Mission 3** | 1 module complémentaire (selon bloc électif) | 6,0 |
| **Relecture & packaging** | Tests, documentation technique, script de déploiement | 1,5 |
| **Total** | | **93,5** |

## Décomposition par profil (TJM)

| Profil | TJM HT | Jours | Total HT |
|---|---|---|---|
| Chef de projet MOE (pilotage, GANTT, Trello, livrables) | 950 € | 7,5 | 7 125,00 € |
| Architecte applicatif senior (modélisation BDD, API Go, i18n, cloud) | 900 € | 4,0 | 3 600,00 € |
| Développeur senior backend (API Go, services métier, authentification) | 800 € | 12,0 | 9 600,00 € |
| Développeur fullstack JavaScript/PHP/frameworks (espaces particuliers, pros, salariés, back-office) | 650 € | 24,0 | 15 600,00 € |
| Intégrateur junior (Stripe, génération PDF, OneSignal, tutoriel overlay, multilingue) | 600 € | 7,0 | 4 200,00 € |
| Ingénieur infrastructure système senior (AD/DNS, GPO, baie de stockage, DMZ mail, backup) | 850 € | 10,0 | 8 500,00 € |
| Ingénieur réseau & sécurité senior (RNSH, clusters HA, VPN, filtrage, certification « Spécialiste bloc Infra ») | 900 € | 13,0 | 11 700,00 € |
| Technicien systèmes & réseaux (GLPI, supervision, 3CX, WiFi, déploiement Docker) | 500 € | 10,0 | 5 000,00 € |
| Spécialiste module complémentaire Mission 3 (selon bloc : Dev / Sciences / Infra) | 700 € | 6,0 | 4 200,00 € |
| **Total devis n°1** | | **93,5 j** | **69 525,00 € HT** |

## Détail de la charge Mission 1 (46 j)

| Tâche | Jours |
|---|---|
| Conception + modélisation (API Go, BDD, multilingue sans service externe) | 3,0 |
| Bootstrap front-end, design system, overlay tutoriel première connexion | 3,0 |
| API Go : authentification/rôles, CRUD, traitement métier (annonces, conteneurs, projets, score) | 8,0 |
| Espace Particuliers (annonces, conteneurs/code-barres, conseils, catalogue, Upcycling Score, planning) | 6,0 |
| Espace Pros/Artisans (abonnements, facturation, annonces, récupération conteneurs, projets) | 5,0 |
| Espace Salariés (formations, planning, conseils/news, forums) | 3,0 |
| Back-office d'administration générale (utilisateurs, acteurs, contrats, financier, validations, notifications, conteneurs) | 6,0 |
| Paiement Stripe | 2,0 |
| Génération automatique de documents PDF + archivage d'un double | 1,5 |
| Notifications push (API OneSignal) | 1,5 |
| Serveur Web personnel (réécriture d'URL, codes d'erreur) + Docker Compose + packaging | 2,5 |
| Tests / recette / qualification des espaces | 4,5 |
| **Total Mission 1** | **46,0** |

## Détail de la charge Mission 2 (32,5 j)

| Tâche | Jours |
|---|---|
| Dimensionnement, documentation technique, devis de l'infrastructure | 1,0 |
| Paris : Active Directory / Domain Controller / DNS (Windows Server 2022) | 2,0 |
| Paris : GPO de base (fond d'écran, lecteurs, mot de passe, installation) | 1,5 |
| Paris : baie de stockage centralisée (TrueNAS) toutes données M1 | 1,5 |
| Paris : DMZ serveur mail Windows (envoi/réception, notifications M1) | 2,0 |
| Paris : déploiement site via containers Docker (Compose) | 1,0 |
| Paris : 3 accès site (externe DMZ, LAN salariés, VPN télétravail) | 1,5 |
| Paris : 8 VLAN (Direction, Marketing, Commercial, RH, Informatique, Hmanager, Regional, DMZ) | 1,0 |
| Paris : cluster HA 2 firewalls OPNSense (imposé) | 2,0 |
| RNSH : routage inter-VLAN, VPN IPSec over GRE, RIP/OSPF, règles de filtrage | 3,0 |
| VPN externe (OpenVPN / WireGuard) | 1,0 |
| GLPI (ticketing + inventaire) + supervision (Zabbix) | 2,5 |
| Backup complet + incrémental (Veeam) | 1,5 |
| Montreuil : VPN site-to-site Paris + client-to-site + cluster HA OPNSense | 2,0 |
| Suisse (RGPD, personnel certifié) : GPO strictes (7), cluster HA 2 PFSense, VPN site-to-site Paris | 2,5 |
| Datacenter : OpenVPN site-to-site certifiés, cluster 2 firewalls (1 PFSense + 1 OPNSense) | 2,0 |
| Box 11e & 13e : OpenVPN site-to-site, WiFi dédié, serveur 3CX + softphones | 2,0 |
| Intégration production, recette globale, démonstration | 2,5 |
| **Total Mission 2** | **32,5** |

> **NB :** la charge Mission 3 (6 j) varie selon le bloc électif retenu (Développement — Sciences/IA — Infra) ; le TJM du profil associé est ajusté en conséquence.

---

# DEVIS N°2 — INFRASTRUCTURE SYSTÈME

Matériel serveur, licences, stockage, hébergement — hors équipements réseau (devis n°3).

## Site de Paris (siège)

| Poste | Qté | PU HT | Total HT |
|---|---|---|---|
| Serveur physique virtualisation (hôte principal, Proxmox/ESXi) | 1 | 3 800 € | 3 800,00 € |
| Serveur physique virtualisation (hôte redondance/failover) | 1 | 3 800 € | 3 800,00 € |
| Licence Windows Server 2022 Standard (DC + DNS + mail DMZ) | 2 | 900 € | 1 800,00 € |
| Baie de stockage TrueNAS 2U 8 × 20 To (toutes données, bases M1) | 2 | 2 800 € | 5 600,00 € |
| VM serveur web / API (containers Docker) | 1 | 1 200 € | 1 200,00 € |
| Licence Veeam Backup & Replication (complet + incrémental) | 1 | 1 500 € | 1 500,00 € |
| Forfait EDR / sécurisation postes parc | 1 | 1 200 € | 1 200,00 € |
| **Sous-total Paris** | | | **18 900,00 €** |

## Site de Montreuil

| Poste | Qté | PU HT | Total HT |
|---|---|---|---|
| Serveur NAS entrepôt (remise à neuf, répliqué vers Paris) | 1 | 2 500 € | 2 500,00 € |
| Licences systèmes + STOCKAGE supplémentaire | 1 | 800 € | 800,00 € |
| **Sous-total Montreuil** | | | **3 300,00 €** |

## Site de Suisse (hors Schengen — RGPD)

| Poste | Qté | PU HT | Total HT |
|---|---|---|---|
| Serveur local (contrôleur de domaine, 4 postes) | 1 | 1 500 € | 1 500,00 € |
| Postes sécurisés + licences Windows (gestion GPO strictes) | 4 | 800 € | 3 200,00 € |
| **Sous-total Suisse** | | | **4 700,00 €** |

## Datacenter

| Poste | Qté | PU HT | Total HT |
|---|---|---|---|
| Serveurs virtualisés (données critiques RGPD) | 2 | 2 000 € | 4 000,00 € |
| Hébergement / colocation (forfait annuel) | 1 | 1 600 € | 1 600,00 € |
| **Sous-total Datacenter** | | | **5 600,00 €** |

## Box 11e et 13e

| Poste | Qté | PU HT | Total HT |
|---|---|---|---|
| Serveur téléphonique 3CX + licences softphones | 2 | 850 € | 1 700,00 € |
| **Sous-total Box** | | | **1 700,00 €** |

## Total devis n°2

| | |
|---|---|
| **Total HT** | **34 200,00 €** |
| TVA 20 % | 6 840,00 € |
| **Total TTC** | **41 040,00 €** |

---

# DEVIS N°3 — INFRASTRUCTURE RÉSEAU

Équipements réseau, pare-feux, routage, WiFi — hors systèmes (devis n°2).

## Site de Paris (siège)

| Poste | Qté | PU HT | Total HT |
|---|---|---|---|
| Cluster HA 2 firewalls **OPNSense** (imposé, non substituable) | 2 | 1 800 € | 3 600,00 € |
| Routeur cœur RNSH (routage inter-VLAN, IPSec) | 1 | 1 800 € | 1 800,00 € |
| Switchs cœur / distribution / accès (VLAN) | 1 | 4 500 € | 4 500,00 € |
| **Sous-total Paris** | | | **9 900,00 €** |

## Site de Montreuil

| Poste | Qté | PU HT | Total HT |
|---|---|---|---|
| Cluster HA 2 firewalls **OPNSense** (imposé) | 2 | 1 500 € | 3 000,00 € |
| Switchs entrepôt | 1 | 1 800 € | 1 800,00 € |
| **Sous-total Montreuil** | | | **4 800,00 €** |

## Site de Suisse (hors Schengen)

| Poste | Qté | PU HT | Total HT |
|---|---|---|---|
| Cluster HA 2 firewalls **PFSense** (imposé, non substituable) | 2 | 1 500 € | 3 000,00 € |
| Routeur / switch local | 1 | 1 200 € | 1 200,00 € |
| **Sous-total Suisse** | | | **4 200,00 €** |

## Datacenter

| Poste | Qté | PU HT | Total HT |
|---|---|---|---|
| Cluster 2 firewalls (1 **PFSense** + 1 **OPNSense**) | 2 | 2 000 € | 4 000,00 € |
| Routeur datacenter | 1 | 1 200 € | 1 200,00 € |
| **Sous-total Datacenter** | | | **5 200,00 €** |

## Box 11e et 13e

| Poste | Qté | PU HT | Total HT |
|---|---|---|---|
| Bornes WiFi gérées (Wifi_11 / Wifi_13) | 1 | 2 500 € | 2 500,00 € |
| Routeur OpenVPN site-to-site | 1 | 900 € | 900,00 € |
| **Sous-total Box** | | | **3 400,00 €** |

## RNSH — interconnexions

| Poste | Qté | PU HT | Total HT |
|---|---|---|---|
| Forfait licences VPN + supervision des liens | 1 | 1 500 € | 1 500,00 € |
| **Sous-total RNSH** | | | **1 500,00 €** |

## Total devis n°3

| | |
|---|---|
| **Total HT** | **29 000,00 €** |
| TVA 20 % | 5 800,00 € |
| **Total TTC** | **34 800,00 €** |

---

# Conditions

- **Validité de l'offre :** 60 jours à compter de la date du devis.
- **Délai de réalisation estimé :** missions 1 et 2 en parallèle, jalons conformes au GANTT livré ; démonstration sur serveur Web déployé (hors localhost, conformément au cahier des charges).
- **Encadrement :** deux chefs de projet (réseau et développement) ; chaque membre participe aux missions 1 et 2 (consigne projet annuel).
- **Gestion de projet :** Trello et GitHub obligatoires, un compte par membre.
- **Matériel de virtualisation :** Eve-NG / GNS3 ; location de serveurs tiers plafonnée à ~40 € (budget groupe) le cas échéant.
- **Conditions de paiement :** 30 % à la commande, 40 % à la livraison partielle (fin M1/M2), 30 % à la recette finale.
- **Horaires :** extrapolés en homme/heure sur la base de 8 h/jour (93,5 j ≈ 748 h).