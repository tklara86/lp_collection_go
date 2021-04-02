const form = document.getElementById('search-albums');

form.addEventListener('submit', (e) => {
    // disable default action
    e.preventDefault();

    let formData = new FormData(form);
    formData.append("csrf_token", "{{.CSRFToken}}")

    // make post request
    fetch('/search-albums', {
        method: 'POST',
    body: formData
    })
        .then(res => res.json())
        .then(data => console.log(data))
        .catch(err => console.error(err));
});

