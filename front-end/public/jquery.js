
    fetch("./public/head.html")
  .then(response => {
    return response.text()
  })
  .then(data => {
    document.querySelector("header").innerHTML = data;
  });

  fetch("./public/footer.html")
  .then(response => {
    return response.text()
  })
  .then(data => {
    document.querySelector("footer").innerHTML = data;
  });


  fetch("./public/nav.html")
  .then(response => {
    return response.text()
  })
  .then(data => {
    document.querySelector("nav").innerHTML = data;
  });
