<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Cat Voter</title>
    <link rel="stylesheet" href="/static/css/style.css">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/swiper@11/swiper-bundle.min.css" />
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css">
</head>
<body>
    <div class="container">
        <main id="mainContent">
            <!-- Navigation Menu -->
            <nav class="nav-menu">
                <button id="votingBtn" class="active">
                    <i class="fas fa-sort"></i>
                    <span>Voting</span>
                </button>
                <button id="breedsBtn">
                    <i class="fas fa-search"></i>
                    <span>Breeds</span>
                </button>
                <button id="favsBtn">
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"></path>
                    </svg>
                    <span>Favs</span>
                </button>
            </nav>

            <!-- Voting Section -->
            <div id="votingSection">
                <img id="catImage" src="" alt="Cat Image">
                <div class="buttons">
                    <button id="favBtn">
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"></path>
                        </svg>
                    </button>
                    <button id="upvoteBtn">
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <path d="M14 9V5a3 3 0 0 0-3-3l-4 9v11h11.28a2 2 0 0 0 2-1.7l1.38-9a2 2 0 0 0-2-2.3zM7 22H4a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h3"></path>
                        </svg>
                    </button>
                    <button id="downvoteBtn">
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <path d="M10 15v4a3 3 0 0 0 3 3l4-9V2H5.72a2 2 0 0 0-2 1.7l-1.38 9a2 2 0 0 0 2 2.3zm7-13h2.67A2.31 2.31 0 0 1 22 4v7a2.31 2.31 0 0 1-2.33 2H17"></path>
                        </svg>
                    </button>
                </div>
            </div>

            <!-- Breeds Section -->
            <div id="breedsSection" style="display:none;">
                <select id="breedSelect">
                    <option value="">Select a breed</option>
                    <!-- Populate this dropdown with breed options dynamically -->
                </select>

                <!-- Swiper -->
                <div class="swiper mySwiper">
                    <div class="swiper-wrapper" id="carouselInner"></div>
                    <!-- Pagination Dots -->
                    <div class="swiper-pagination"></div>
                </div>

                <!-- Breed Info -->
                <div id="breedInfo"></div>
            </div>

            <!-- Favourites Section -->
            <div id="favsSection" class="section">
                <div id="viewButtons">
                    <button id="gridViewBtn">
                        <i class="fas fa-th-large"></i>
                    </button>
                    <button id="listViewBtn">
                        <i class="fas fa-list"></i>
                    </button>
                </div>
                <div id="favsContainer" class="favs-grid">
                    <!-- Favourite images will be displayed here -->
                </div>
            </div>
        </main>
    </div>

    <!-- Scripts -->
    <script src="/static/js/app.js"></script>
    <script src="/static/js/voting.js"></script>
    <script src="/static/js/breeds.js"></script>
    <script src="/static/js/favourite.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/swiper@11/swiper-bundle.min.js"></script>
</body>
</html>
