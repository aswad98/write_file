const data = {
    file_name  : document.getElementById("filename").value,
    statment : document.getElementById("filedata").value
}
const form = document.getElementById("submit").addEventListener("click" , async function(e){
e.preventDefault();
const response = await fetch('http://localhost:1323/writefile' , {method : 'POST' , body :JSON.stringify(data) , headers: {
    'Content-Type': 'application/json'
}});
const myJson = await response.json();
console.log(myJson)
})
console.log(data);