// favorites.js

document.addEventListener('DOMContentLoaded', function () {
    const favsContainer = document.getElementById('favsContainer');
    const gridViewBtn = document.getElementById('gridViewBtn');
    const listViewBtn = document.getElementById('listViewBtn');
    let currentUser = 'user-123';
    
    function showFavourites(view) {
        fetch(`/api/favourites?sub_id=${currentUser}&limit=10&page=0`)
            .then(response => response.json())
            .then(favs => {
                // console.log(favs);
                
                favsContainer.innerHTML = '';
                favsContainer.className = view === 'grid' ? 'favs-grid' : 'favs-list';
                favs.forEach(fav => {
                    const img = document.createElement('img');
                    img.src = fav.image.url;
                    img.alt = 'Favourite Cat';
                    img.dataset.favouriteId = fav.id;

                    const deleteBtn = document.createElement('button');
                    deleteBtn.textContent = 'Remove';
                    deleteBtn.classList.add('remove-btn');
                    deleteBtn.onclick = () => removeFavourite(fav.id);

                    const container = document.createElement('div');
                    container.classList.add('image-container');
                    container.appendChild(img);
                    container.appendChild(deleteBtn);
                    favsContainer.appendChild(container);
                });
            })
            .catch(error => console.error('Error fetching favourites:', error));
    }

    function removeFavourite(favouriteId) {
        fetch(`/api/favourites/${favouriteId}`, {
            method: 'DELETE'
        })
            .then(response => response.json())
            .then(data => {
                if (data.message === 'SUCCESS') {
                    alert('Removed from favourites!');
                    showFavourites(favsContainer.className.includes('grid') ? 'grid' : 'list');
                } else {
                    alert('Error removing from favourites');
                }
            })
            .catch(error => console.error('Error removing favourite:', error));
    }

    // Event listeners to toggle views
    gridViewBtn.addEventListener('click', () => showFavourites('grid'));
    listViewBtn.addEventListener('click', () => showFavourites('list'));

    // Initial load
    if (favsContainer.style.display !== 'none') {
        showFavourites('grid');
    }
});
