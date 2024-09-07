
function rec(elm) {
	if(elm.nodeName==="CODE"){
    if(elm.parentNode.nodeName != "PRE" ){
      elm.className = styleMap["CODE"] ;
    }
    else {
      return ; 
    }
  }

  if(elm.nodeName==="BLOCKQUOTE"){
    elm.className = styleMap["BLOCKQUOTE"]
  }

	if(elm.nodeName==="LI"){
		if(elm.parentNode.nodeName==="OL"){
      elm.className =styleMap["OL_LI"]; 
    }
	  else{
		  elm.className = styleMap["UL_LI"] ;
	  }
  }

  if (elm.children.length === 0) {
    if (elm.nodeName === "LI") {
      elm.className = styleMap["OL_LI"];
    } else {
      elm.className = styleMap[elm.nodeName] || ""; 
    }
  } else {
    for (let i = 0; i < elm.children.length; i++) {
      rec(elm.children[i]);
    }
  }
}

function renderMarkdown() {
    const input =document.getElementById('thing').innerText;
    const result = marked.parse(input);
    document.getElementById('markdow').innerHTML = result;
    const markdown = document.getElementById('markdow') ;
    rec(markdown) ;

}

renderMarkdown();
    
    
hljs.highlightAll()
