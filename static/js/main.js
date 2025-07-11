const btn = document.querySelector(".generate button");

const btnHandler = (event) => {
    const min = parseInt(document.getElementById("min").value);
    const max = parseInt(document.getElementById("max").value);

    const result = document.getElementById("result");

    if(min > max) {
        result.className = "result-err";
        result.innerHTML = "Min value cannot be greater than max value";
        return;
    }

    if(min == max) {
        result.className = "result-err";
        result.innerHTML = "Min value cannot be equal to or greater than max value";
        return;
    }

    fetch(`/api/randomNumber?min=${min}&max=${max}`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json"
        },
    })
    .then(response => response.json())
    .then(data => {
        value = data.value;

        result.className = "result-succ";
        result.innerHTML = value;
    })
    .catch(error => {
        console.log(error);
    });

}

btn.addEventListener("click", btnHandler);