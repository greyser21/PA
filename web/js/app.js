document.addEventListener("DOMContentLoaded", () => {
    initDates();
    initDemoButtons();
    initPage();
});

function initDates() {
    const dateElements = document.querySelectorAll("[data-current-date]");
    const now = new Date();
    const formattedDate = now.toLocaleDateString("fr-FR", {
        weekday: "long", day: "numeric", month: "long", year: "numeric"
    });
    dateElements.forEach(el => el.textContent = formattedDate);
}

function initDemoButtons() {
    document.querySelectorAll("[data-demo-action]").forEach(button => {
        button.addEventListener("click", () => {
            showToast(button.dataset.demoAction || "Fonctionnalité en cours de développement.");
        });
    });
}

function showToast(message) {
    let toast = document.querySelector(".toast");
    if (!toast) {
        toast = document.createElement("div");
        toast.className = "toast";
        document.body.appendChild(toast);
    }
    toast.textContent = message;
    toast.classList.add("visible");
    clearTimeout(window.upcycleToastTimer);
    window.upcycleToastTimer = setTimeout(() => toast.classList.remove("visible"), 2500);
}

async function apiGet(path) {
    const res = await fetch(path);
    if (!res.ok) throw new Error("API error " + res.status);
    return res.json();
}

async function apiPost(path, body) {
    const res = await fetch(path, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(body)
    });
    const data = await res.json().catch(() => ({}));
    if (!res.ok) throw new Error(data.error || ("API error " + res.status));
    return data;
}

function renderStats(containerId, stats) {
    const el = document.getElementById(containerId);
    if (!el) return;
    el.innerHTML = stats.map(s => `
        <article class="stat-card"><span class="label">${escapeHtml(s.label)}</span><strong>${escapeHtml(String(s.value))}</strong></article>
    `).join("");
}

function escapeHtml(str) {
    return str.replace(/[&<>"']/g, c => ({'&':'&amp;','<':'&lt;','>':'&gt;','"':'&quot;',"'":'&#39;'}[c]));
}

function initPage() {
    const page = document.body.dataset.page;
    if (page === "user") loadUser();
    if (page === "pro") loadPro();
    if (page === "employee") loadEmployee();
    if (page === "admin") loadAdmin();
}

async function loadUser() {
    try {
        const data = await apiGet("/api/user/dashboard");
        const nameEl = document.getElementById("user-name");
        if (nameEl) nameEl.textContent = "utilisateur";
        renderStats("user-stats", data.stats || []);
        renderList("user-announcements", (data.announcements || []).map(a => ({
            title: a.title,
            subtitle: `${a.type}${a.location ? " • " + a.location : ""}${a.price ? " • " + a.price : ""}`,
            badge: a.status
        })), "Aucune annonce pour le moment.");
        renderList("user-activities", (data.activities || []).map(a => ({
            title: a.title, subtitle: a.date, badge: a.status
        })), "Aucune activité planifiée.");
    } catch (e) {
        showToast("Impossible de charger le tableau de bord utilisateur.");
    }
    const form = document.getElementById("announcement-form");
    if (form && !form.dataset.bound) {
        form.dataset.bound = "1";
        form.addEventListener("submit", async (ev) => {
            ev.preventDefault();
            const fd = new FormData(form);
            const payload = {
                title: fd.get("title"), type: String(fd.get("type") || "").toLowerCase(),
                location: fd.get("location") || "", price: fd.get("price") || "", status: "en_attente"
            };
            const fb = document.getElementById("announcement-feedback");
            try {
                await apiPost("/api/announcements", payload);
                if (fb) fb.textContent = "Annonce envoyée, en attente de validation administrative.";
                form.reset();
                showToast("Annonce créée (en attente de validation).");
                loadUser();
            } catch (err) {
                if (fb) fb.textContent = "Erreur : " + err.message;
            }
        });
    }
    const dep = document.getElementById("btn-deposit");
    if (dep && !dep.dataset.bound) {
        dep.dataset.bound = "1";
        dep.addEventListener("click", async () => {
            try {
                const r = await apiPost("/api/deposits", { item: "Objet à déposer", location: "Paris", status: "en_attente" });
                showToast("Demande de dépôt créée (" + (r.barcode || r.openCode || "en attente") + ").");
                loadUser();
            } catch (e) { showToast("Erreur lors de la demande de dépôt."); }
        });
    }
}

async function loadPro() {
    try {
        const data = await apiGet("/api/pro/dashboard");
        renderStats("pro-stats", data.stats || []);
        renderList("pro-items", (data.items || []).map(i => ({
            title: i.material,
            subtitle: `Conteneur ${i.container} • ${i.location}`,
            badge: i.status
        })), "Aucun objet à récupérer.");
        renderList("pro-projects", (data.projects || []).map(p => ({
            title: p.title,
            subtitle: `${p.details || ""}${p.progress != null ? " • " + p.progress + " % terminé" : ""}`,
            badge: p.status
        })), "Aucun projet.");
    } catch (e) {
        showToast("Impossible de charger le tableau de bord professionnel.");
    }
}

async function loadEmployee() {
    try {
        const data = await apiGet("/api/employee/dashboard");
        renderStats("employee-stats", data.stats || []);
        renderList("employee-activities", (data.activities || []).map(a => ({
            title: a.title, subtitle: a.date, badge: a.status
        })), "Aucune activité.");
        renderList("employee-tasks", (data.tasks || []).map(t => ({
            title: t.title, subtitle: t.details, badge: t.status
        })), "Rien à traiter.");
    } catch (e) {
        showToast("Impossible de charger l'espace salarié.");
    }
}

async function loadAdmin() {
    try {
        const data = await apiGet("/api/admin/dashboard");
        renderStats("admin-stats", data.stats || []);
        renderList("admin-actions", (data.actions || []).map(a => ({
            title: a.title, subtitle: a.details, badge: a.status
        })), "Aucune action prioritaire.");
        renderList("admin-activity", (data.activity || []).map(a => ({
            title: a.title, subtitle: a.details, badge: a.when
        })), "Aucune activité récente.");
    } catch (e) {
        showToast("Impossible de charger le back-office.");
    }
}

function renderList(containerId, items, emptyMsg) {
    const el = document.getElementById(containerId);
    if (!el) return;
    if (!items || items.length === 0) {
        el.innerHTML = `<div class="list-item"><small>${escapeHtml(emptyMsg)}</small></div>`;
        return;
    }
    el.innerHTML = items.map(it => `
        <div class="list-item">
            <div><strong>${escapeHtml(it.title || "")}</strong><small>${escapeHtml(it.subtitle || "")}</small></div>
            <span class="badge">${escapeHtml(it.badge || "")}</span>
        </div>
    `).join("");
}
