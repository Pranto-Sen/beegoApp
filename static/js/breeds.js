// breeds.js

document.addEventListener('DOMContentLoaded', function () {
    const breedSelect = document.getElementById('breedSelect');
    const breedInfo = document.getElementById('breedInfo');
    let swiper;

    function fetchBreeds() {
        fetch('/api/breeds')
            .then(response => response.json())
            .then(breeds => {
                breedSelect.innerHTML = '<option value="">Select a breed</option>';
                breeds.forEach(breed => {
                    const option = document.createElement('option');
                    option.value = breed.id;
                    option.textContent = breed.name;
                    breedSelect.appendChild(option);
                });

                if (breeds.length > 0) {
                    breedSelect.value = breeds[0].id;
                    breedSelect.dispatchEvent(new Event('change'));
                }
            })
            .catch(error => console.error('Error fetching breeds:', error));
    }

    breedSelect.addEventListener('change', (e) => {
        const selectedBreed = e.target.value;
        if (selectedBreed) {
            fetch(`https://api.thecatapi.com/v1/breeds/${selectedBreed}`)
                .then(response => response.json())
                .then(breed => {
                    const breedId = breed.id;
                    fetch(`https://api.thecatapi.com/v1/images/search?breed_ids=${breedId}&limit=5`)
                        .then(response => response.json())
                        .then(images => {
                            const carouselInner = document.getElementById('carouselInner');
                            carouselInner.innerHTML = '';

                            images.forEach(image => {
                                if (image.url) {
                                    carouselInner.innerHTML += `
                                        <div class="swiper-slide">
                                            <img src="${image.url}" alt="${breed.name}">
                                        </div>
                                    `;
                                } else {
                                    console.error("No image URL found for this image object:", image);
                                }
                            });

                            if (swiper) {
                                swiper.destroy();
                            }
                            swiper = new Swiper(".mySwiper", {
                                spaceBetween: 30,
                                pagination: {
                                    el: ".swiper-pagination",
                                    clickable: true,
                                },
                                autoplay: {
                                    delay: 3000,
                                    disableOnInteraction: false,
                                },
                            });

                            breedInfo.innerHTML = `
                                <div class="breed-info">
                                    <h2>${breed.name}</h2>
                                    <p class="origin">(${breed.origin})</p>
                                    <p class="id">${breed.id}</p>
                                </div>
                                <p>${breed.description}</p>
                                <a class="breed-wiki" href="${breed.wikipedia_url}">WIKIPEDIA</a>
                            `;
                        })
                        .catch(error => console.error('Error fetching images:', error));
                })
                .catch(error => console.error('Error fetching breed:', error));
        } else {
            document.getElementById('carouselInner').innerHTML = '';
            breedInfo.innerHTML = '';
            if (swiper) {
                swiper.destroy();
                swiper = null;
            }
        }
    });

    // Initial load
    if (breedSelect.style.display !== 'none') {
        fetchBreeds();
    }
});
