/*function copyClipboard() {
    let textCopied = document.getElementById("autoExpand")

    textCopied.select();
    textCopied.setSelectionRange(0, 99999);*/
/* Mobile version */
/*document.execCommand('copy');

    alert("Texte copié !")
}*/

function copyClipboard() {
    const copyText = document.querySelector("pre").textContent;
    const textArea = document.createElement('textarea');
    textArea.textContent = copyText;
    document.body.append(textArea);
    textArea.select();
    document.execCommand("copy");
}

//document.getElementById('button').addEventListener('click', copyClipboard);