document.addEventListener("DOMContentLoaded", () => {
     // Toggle menu for mobile view
     const menuToggle = document.querySelector(".menu-toggle");
     const navLinks = document.querySelector(".nav-links");
   
     menuToggle.addEventListener("click", () => {
       navLinks.classList.toggle("active");
       menuToggle.classList.toggle("open");
     });
   
     // Set the active class based on the current URL
     const currentPath = window.location.pathname;
     const links = document.querySelectorAll(".nav-links a");
   
     links.forEach(link => {
       // Check if the link's href matches the current path
       if (link.getAttribute("href") === currentPath.split("/").pop()) {
         link.classList.add("active");
       } else {
         link.classList.remove("active");
       }
     });
   });
   