/* ============================================================
   CLAUSE CLI – Navigation
   ============================================================ */
document.addEventListener('DOMContentLoaded', function () {
    const mobileToggle = document.querySelector('.mobile-menu-toggle');
    const mobileNav = document.querySelector('.mobile-nav');

    if (mobileToggle && mobileNav) {
        mobileToggle.addEventListener('click', function () {
            mobileToggle.classList.toggle('open');
            mobileNav.classList.toggle('open');
            document.body.style.overflow = mobileNav.classList.contains('open') ? 'hidden' : '';
        });

        // Close on link click
        mobileNav.querySelectorAll('a').forEach(function (link) {
            link.addEventListener('click', function () {
                mobileToggle.classList.remove('open');
                mobileNav.classList.remove('open');
                document.body.style.overflow = '';
            });
        });

        // Close on Escape
        document.addEventListener('keydown', function (e) {
            if (e.key === 'Escape' && mobileNav.classList.contains('open')) {
                mobileToggle.classList.remove('open');
                mobileNav.classList.remove('open');
                document.body.style.overflow = '';
            }
        });
    }

    // Active link highlighting
    const currentPath = window.location.pathname.replace(/\/$/, '') || '/';
    document.querySelectorAll('.main-nav a, .doc-sidebar-links a').forEach(function (a) {
        const href = a.getAttribute('href');
        if (href && (currentPath.endsWith(href.replace(/\/$/, '')) || href === currentPath)) {
            a.classList.add('active');
        }
    });
});
