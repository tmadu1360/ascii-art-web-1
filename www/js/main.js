function copyClipboard() {
    const copyText = document.querySelector("pre").textContent;
    const textArea = document.createElement('textarea');
    textArea.textContent = copyText;
    document.body.append(textArea);
    textArea.select();
    document.execCommand("copy");
    alert("Texte copié !");
}

//document.getElementById('button').addEventListener('click', copyClipboard);