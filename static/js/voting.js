document.addEventListener('DOMContentLoaded', function () {
    const votingSection = document.getElementById('votingSection');
    const catImage = document.getElementById('catImage');
    const upvoteBtn = document.getElementById('upvoteBtn');
    const downvoteBtn = document.getElementById('downvoteBtn');
    const favBtn = document.getElementById('favBtn');

    let currentCat = null;
    let currentUser = 'user-123';

    function fetchCat() {
        fetch('/api/cat')
            .then(response => response.json())
            .then(data => {
                currentCat = data;
                catImage.src = data.url;
            })
            .catch(error => console.error('Error fetching cat:', error));
    }

    function addFavourite(imageId) {
        const rawBody = JSON.stringify({
            "image_id": imageId,
            "sub_id": currentUser
        });

        fetch('/api/favourites', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: rawBody
        })
            .then(response => response.json())
            .then(data => {
                if (data.error) {
                    console.error('Error:', data.error);
                } else {
                    alert("Successfully added to favourites");
                    console.log('Success:', data);
                }
            })
            .catch(error => {
                console.error('Fetch error:', error);
            });
    }

    upvoteBtn.addEventListener('click', fetchCat);
    downvoteBtn.addEventListener('click', fetchCat);

    favBtn.addEventListener('click', () => {
        if (currentCat) {
           
            addFavourite(currentCat.id);
        }
        fetchCat();
    });

    if (votingSection.style.display !== 'none') {
        fetchCat();
    }
});
