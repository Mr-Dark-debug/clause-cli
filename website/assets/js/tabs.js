/* ============================================================
   CLAUSE CLI – Tabs Component
   ============================================================ */
document.addEventListener('DOMContentLoaded', function () {
    document.querySelectorAll('.tabs').forEach(function (tabsContainer) {
        const buttons = tabsContainer.querySelectorAll('.tab');
        const panels = tabsContainer.querySelectorAll('.tab-panel');

        buttons.forEach(function (btn) {
            btn.addEventListener('click', function () {
                const target = btn.getAttribute('data-tab');

                // Deactivate all
                buttons.forEach(function (b) { b.classList.remove('active'); });
                panels.forEach(function (p) { p.classList.remove('active'); });

                // Activate clicked
                btn.classList.add('active');
                var panel = tabsContainer.querySelector('[data-panel="' + target + '"]');
                if (panel) panel.classList.add('active');
            });
        });
    });
});
