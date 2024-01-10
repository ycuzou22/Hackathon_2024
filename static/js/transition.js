const card = document.getElementById('card');
const iconContainer = document.getElementById('iconContainer');
const moonIcon = document.getElementById('moonIcon');
const sunIcon = document.getElementById('sunIcon');

let isFlipped = false;

iconContainer.addEventListener('click', () => {
    card.style.transform = isFlipped ? 'rotateY(0deg)' : 'rotateY(180deg)';
    moonIcon.style.display = isFlipped ? 'inline' : 'none';
    sunIcon.style.display = isFlipped ? 'none' : 'inline';
    isFlipped = !isFlipped;
});
