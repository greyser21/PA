document.addEventListener("DOMContentLoaded", () => {
    // Date affichée dans les tableaux de bord
    const dateElements = document.querySelectorAll("[data-current-date]");
    const now = new Date();

    const formattedDate = now.toLocaleDateString("fr-FR", {
        weekday: "long",
        day: "numeric",
        month: "long",
        year: "numeric"
    });

    dateElements.forEach(element => {
        element.textContent = formattedDate;
    });

    // Petite interaction générique pour les boutons "démo"
    document.querySelectorAll("[data-demo-action]").forEach(button => {
        button.addEventListener("click", () => {
            const message = button.dataset.demoAction || "Fonctionnalité en cours de développement.";
            showToast(message);
        });
    });
});

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
    window.upcycleToastTimer = setTimeout(() => {
        toast.classList.remove("visible");
    }, 2500);
}
