# Project: Adoption-OS
**Goal:** A self-hosted cat adoption platform running locally for "I Dream of Home Rescue."

## 1. Hardware & Storage Architecture
* **Server:** Raspberry Pi 5 (8GB RAM).
* **OS & Code:** Hosted on **SD Card** (Standard boot partition).
* **Primary Data Storage:** **1TB NVMe Drive** (PCIe attached).
    * *Purpose A:* **PostgreSQL Database** (The actual data directory must be moved here).
    * *Purpose B:* **Pet Photos/Assets** (High-res images).
    * *Implication:* All persistent data lives here to prevent SD card wear and ensure high I/O performance.

## 2. Technology Stack
* **Backend:** Go (Golang) standard library + `net/http`.
* **Frontend:** Vue 3 + Vite + **Pinia** (State Management).
* **Database:** PostgreSQL (local).
    * *Schema:* Managed via migrations (`backend/db/migrations`).
    * *Data:* `pets` table uses JSONB for complex traits.
* **Infrastructure (Planned):** **Nginx**.
    * *Role:* Reverse Proxy listening on Port 80.
    * *Strategy:* * `/api/*` -> Proxies to Go Backend (:8080).
        * `/images/*` -> Serves directly from NVMe (High performance).

## 3. Development Focus (Strict Phasing)
**Current Focus: Phase 1**
* **Phase 1: The Foundation (MVP) (Immediate):**
    * **Public Website:** Responsive Home, About, Adopt, Volunteer, and Donate pages.
        * * Display "Available Pets" and "Spotlight" lists using Pinia stores.
        * * Implement "Pet Surrender" form and logic.
        * * Implement "Schedule Visit" form and logic.
    * **Digital Intake:** Smart SurrenderPet forms with risk logic.
    * **Volunteer Portal:** Application with "Teen vs. Adult" logic.
    * **Backend Core:** Go server & PostgreSQL setup.
        * * Implement Backend API endpoints.
        * * Implement email notifications for "Pet Adoption", "Volunteer Signup", and "Pet Surrender".
        * * Ensure database seeding (`make seed`) is reliable.
    * **Admin Shell:** Basic secure login and navigation.

**Future Phases (Do Not Implement Yet):**
* **Phase 2: Revenue & Reach**
    * **Stripe Foundation:** Process Adoption Fees & One-Time Donations.
    * **Pet Record Management:** Admin CRUD for pet profiles.
    * **Manual Donation & Cash Handling:** Log offline donations (cash/check).
    * **Digital Kennel Cards (E-Ink V1):** Raspberry Pi integration.
    * **Smart QR Codes:** Auto-generate unique pet profile links.

* **Phase 3: Onboarding & Communication**
    * **Social Media Share Generator:** Create "Adopt Me" graphics.
    * **Automated Adopter Onboarding:** Triggered emails with care guides.
    * **Foster-to-Adopt Fast-Track:** One-click adoption for fosters.
    * **SMS Status Integration:** Application status text alerts.
    * **Post-Adoption Surveys:** Automated 3-3-3 check-in emails.

* **Phase 4: Engagement Campaigns**
    * **External API Sync:** Nightly export to Petfinder/RescueGroups.
    * **"Name the Litter":** Paid voting module for naming pets.
    * **Digital Raffle Manager:** Sell tickets & pick winners via Stripe.
    * **Volunteer Leaderboard:** Gamified tracking of volunteer impact.
    * **"Happy Tails":** Public success story submission form.

* **Phase 5: Mobile & Field Operations (PWA)**
    * **Offline "Field Mode":** Service worker caching for intake.
    * **Push Notifications:** Browser alerts for urgent needs.
    * **"Add to Home Screen":** Native-app feel configuration.
    * **Mobile Image Compression:** Client-side resizing for faster uploads.
    * **Geolocation Intake:** GPS tagging for stray intakes.

* **Phase 6: Financial Stability**
    * **Monthly Subscriptions:** Recurring donor engine.
    * **Pet Sponsorships:** Link payments to specific animals.
    * **"Gotcha Day" Engine:** Anniversary donation emails.
    * **Donor Portal:** Self-service receipt & plan management.
    * **Tribute Giving:** "In Memory Of" donation flow.
    * **Legacy Giving Portal:** Planned giving & bequests section.

* **Phase 7: Advanced Hardware**
    * **Intake Kiosk:** Tablet-optimized intake UI.
    * **Vet Dashboard:** "Airport Screen" for medical schedules.
    * **Device Health Monitor:** Admin view for IoT uptime.
    * **NFC Support:** Tap-to-view medical records on collars.
    * **Smart Supply Buttons:** IoT buttons for inventory alerts.

* **Phase 8: Data & Inventory**
    * **Weight Visualization:** Growth charts for neonates.
    * **Inventory System:** "Supply Closet" tracking & alerts.
    * **Medication Logs:** Strict controlled substance auditing.
    * **Volunteer Hours:** Clock-in/out tracking for grants.
    * **Incident Reporting:** Liability forms for bites/injuries.

* **Phase 9: Matchmaking & Discovery**
    * **"Find My Match" Quiz:** Lifestyle-based pet filtering.
    * **"Pet Alerts":** Saved search notifications for users.
    * **Lifestyle Badges:** Auto-tagging (e.g., "Apartment Friendly").
    * **Compare View:** Side-by-side pet comparison tool.
    * **Video Embeds:** YouTube/TikTok support in profiles.

* **Phase 10: Community Resources**
    * **Lost & Found Generator:** Public flyer creation tool.
    * **Community Adoption Events Calendar:** Manage off-site adoption events & volunteer shifts.
    * **Community Calendar:** Public view of clinics/fundraisers.
    * **Partner Directory:** Vets & trainers listing.
    * **Newsletter Builder:** Drag-and-drop email tool.

* **Phase 11: Specialized Care & Crisis**
    * **"SOS" Broadcast:** Emergency SMS to skilled volunteers.
    * **Fospice Tracker:** Quality of Life logs for hospice pets.
    * **Smart Visiting Hours:** Real-time "Open/Closed" widget.
    * **Isolation Management:** Contagion tracking workflows.
    * **Disaster Mode:** Crisis homepage takeover switch.

* **Phase 12: Compliance, Security & Standards**
    * **SOC2 Readiness:** Immutable Audit Logs & Centralized Security Monitoring.
    * **GDPR/CCPA Privacy Suite:** "Right to be Forgotten" engine, Data Export (JSON), & Cookie Consent.
    * **Strict RBAC Enforcement:** Granular permissioning for Staff vs. Admins (prep for SaaS).
    * **Accessibility Deep-Dive:** WCAG 2.1 AA Audit & Screen Reader optimization.
    * **Penetration Testing:** Third-party security audit & vulnerability patching.
    * **Data Retention:** Automated PII deletion policies.

* **Phase 13: SaaS Multi-Tenancy**
    * **Tenant Logic:** Database sharding by organization_id.
    * **Usage Metering:** Track storage/SMS limits per org.
    * **Stripe Billing Portal:** SaaS subscription management.
    * **Feature Gating:** Code-level permission checks.
    * **Setup Wizard:** Automated onboarding for new shelters.

* **Phase 14: Advanced Scoring Logic (Part 1 - Adopters)**
    * **Adopter Scoring Engine:** Algorithm to score applications.
    * **Risk Flags:** Auto-highlight "Red Flags."
    * **Application Sorting:** Default sort by "Score."
    * **Auto-Decline:** Logic to auto-reject unqualified apps.
    * **Reviewer Assignment:** Auto-assign based on workload.

* **Phase 15: Advanced Scoring Logic (Part 2 - Workforce)**
    * **Volunteer Scoring System:** Rank applicants by skill/availability.
    * **Foster Scoring System:** Rate potential fosters.
    * **Reliability Index:** Track "Show-up Rate."
    * **Capacity Planning:** Forecast volunteer needs.
    * **Skill-Gap Analysis:** Identify missing volunteer skills.

* **Phase 16: Waitlist & Demand Management**
    * **Waitlist System:** Users sign up for breeds.
    * **Auto-Notification:** Email waitlist on intake match.
    * **Priority Access:** Early access for Sponsors.
    * **Demand Heatmap:** View requested breeds vs. intake.
    * **"Coming Soon" Teasers:** Blurred profiles for medical hold.

* **Phase 17: Medical Intelligence**
    * **Medical Reminder System:** Alerts for boosters/meds.
    * **Treatment Plans:** Templated medical workflows.
    * **Vet Partner Portal:** External vet upload access.
    * **Cost-per-Animal Tracking:** Medical spend per pet.
    * **Symptom Checker:** Triage tool for fosters.

* **Phase 18: Operational Automation**
    * **Task Automation:** IFTTT logic for shelter tasks.
    * **Document Generator:** Auto-fill PDF contracts.
    * **Shift Reminders:** SMS/Email shift reminders.
    * **Lost & Found Board:** Internal matching logic.
    * **Key Management:** Digital key checkout log.

* **Phase 19: Social & Marketing Automation**
    * **Social Auto-Poster:** Auto-draft posts for new pets.
    * **Story Highlight Generator:** "Weekly Recap" video creator.
    * **Ad Integration:** Boost Facebook posts from admin.
    * **Influencer Portal:** Media kits for promoters.
    * **UTM Tracking:** Track adoption sources.

* **Phase 20: Foster Experience 2.0**
    * **Foster Portal Dashboard:** Central hub for fosters.
    * **Supply Request Flow:** One-click reorders.
    * **Foster-to-Foster Chat:** Secure advice channels.
    * **Vacation Coverage:** "I'm Away" toggle.
    * **Photo Uploader:** Easy mobile photo submission.

* **Phase 21: Data Analytics & Reporting**
    * **Adoption Funnel Analytics:** Applicant drop-off analysis.
    * **Length-of-Stay Reports:** Trends by breed/age.
    * **Donor Lifetime Value (LTV):** High-value supporter ID.
    * **Impact Reports:** "Year in Review" generator.
    * **Heatmaps:** Geo-analysis of adopters/intakes.

* **Phase 22: Integrations (The "Walled Garden")**
    * **Microchip Registry Sync:** API integration for registration.
    * **Internal Accounting Ledger:** Track income/expenses securely.
    * **Internal Email System:** Marketing campaigns without 3rd parties.
    * **Internal Chat System:** Secure team communication.
    * **Internal Calendar System:** Shifts/Events management.

* **Phase 23: AI & Machine Learning**
    * **AI Bio Generator:** Creative text from tags.
    * **Photo Quality Scorer:** Rate photos for adoption appeal.
    * **Churn Prediction:** Identify at-risk volunteers.
    * **Breed Identification:** Computer vision breed estimates.
    * **Sentiment Analysis:** Analyze feedback for issues.

* **Phase 24: Enterprise Expansion**
    * **Multi-Location Support:** Manage multiple sites.
    * **Franchise Model:** Parent/Child org structures.
    * **API Access:** Open API for enterprise tools.
    * **White Labeling:** Branding removal.
    * **Global Localization:** Multi-language/currency support.

## 4. Security & Conventions
* **Secrets:** Load via `.env`.
* **Hardware Security:** LAN-only trust model for now.
* **Validation:** All inputs must be sanitized on the backend.