const transition = document.querySelector(".transition");
setTimeout(function () {
    transition.classList.remove("active");
}, 400);

const liens = document.querySelectorAll('nav a');

for (i = 0; i < liens.length; i++) {
    let lien = liens[i];

    lien.addEventListener('click', function (event) {
        event.preventDefault();
        transition.classList.add('active');

        // Récupérer la nouvelle URL de la page cible
        let nouvellePage = '';

        // Vérifier quel lien a été cliqué et attribuer la nouvelle URL en conséquence
        if (this.getAttribute('href') === 'home.html') {
            nouvellePage = 'reverse_home.html'; // Lien inverse de home.html
        } else if (this.getAttribute('href') === 'reverse_home.html') {
            nouvellePage = 'home.html'; // Lien inverse de reverse_home.html
        }

        setTimeout(function () {
            window.location.href = nouvellePage; // Rediriger vers la nouvelle page après l'animation
        }, 400);
    });
}
