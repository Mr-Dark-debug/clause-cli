/* ============================================================
   CLAUSE CLI – Copy Code
   ============================================================ */
document.addEventListener('DOMContentLoaded', function () {
    document.querySelectorAll('.copy-button').forEach(function (btn) {
        btn.addEventListener('click', function () {
            const block = btn.closest('.code-block');
            const code = block ? block.querySelector('code') : null;
            if (!code) return;

            const text = code.textContent;
            navigator.clipboard.writeText(text).then(function () {
                btn.classList.add('copied');
                const original = btn.innerHTML;
                btn.innerHTML = '<span>✓ Copied!</span>';
                setTimeout(function () {
                    btn.classList.remove('copied');
                    btn.innerHTML = original;
                }, 2000);
            }).catch(function () {
                // Fallback
                const ta = document.createElement('textarea');
                ta.value = text;
                ta.style.position = 'fixed';
                ta.style.opacity = '0';
                document.body.appendChild(ta);
                ta.select();
                document.execCommand('copy');
                document.body.removeChild(ta);
                btn.classList.add('copied');
                const original = btn.innerHTML;
                btn.innerHTML = '<span>✓ Copied!</span>';
                setTimeout(function () {
                    btn.classList.remove('copied');
                    btn.innerHTML = original;
                }, 2000);
            });
        });
    });
});
