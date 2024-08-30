
document.addEventListener('DOMContentLoaded', function () {
    const votingBtn = document.getElementById('votingBtn');
    const breedsBtn = document.getElementById('breedsBtn');
    const favsBtn = document.getElementById('favsBtn');
    const votingSection = document.getElementById('votingSection');
    const breedsSection = document.getElementById('breedsSection');
    const favsSection = document.getElementById('favsSection');

    function showSection(section, path) {
        votingSection.style.display = 'none';
        breedsSection.style.display = 'none';
        favsSection.style.display = 'none';
        section.style.display = 'block';
        history.pushState(null, '', path); // Update the route without refreshing
    }

    function setActiveButton(button) {
        [votingBtn, breedsBtn, favsBtn].forEach(btn => btn.classList.remove('active'));
        button.classList.add('active');
    }

    votingBtn.addEventListener('click', () => {
        setActiveButton(votingBtn);
        showSection(votingSection, '/voting');
    });

    breedsBtn.addEventListener('click', () => {
        setActiveButton(breedsBtn);
        showSection(breedsSection, '/breed');
    });

    favsBtn.addEventListener('click', () => {
        setActiveButton(favsBtn);
        showSection(favsSection, '/favourite');
    });

    // Show the correct section based on the current URL path
    const path = window.location.pathname;
    if (path === '/breed') {
        setActiveButton(breedsBtn);
        showSection(breedsSection, '/breed');
    } else if (path === '/favourite') {
        setActiveButton(favsBtn);
        showSection(favsSection, '/favourite');
    } else {
        setActiveButton(votingBtn);
        showSection(votingSection, '/voting');
    }
});
