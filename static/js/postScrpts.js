
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
  if(elm.nodeName === "P"){
    elm.className = styleMap["P"]
  }
  if(elm.nodeName==="P" && elm.children.length>0){
    const father = elm.parentNode ; 
    if (father.nodeName == "LI"){
      for(const child of elm.children){
        if(child.nodeName === "STRONG"){
          child.className = styleMap["STRONG"]; 
          father.insertBefore(child,father.firstChild)
          father.removeChild(elm)
          return
        }
      }
    }
  }
	if(elm.nodeName==="LI"){
    //console.log(elm)
    //for (let i = 0; i < elm.children.length; i++) {
    //  if(elm.children[i].nodeName==="P"){
    //    let father =elm.children[i] ; 
    //    if(father.children[0].nodeName === "STRONG"){
    //      console.log(father)
    //      child = father.children[0] ;
    //      father.parentNode.insertBefore(child ,father);
    //
    //    }
    //  } 
    //}
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
            // Get the Markdown content
            const input = document.getElementById('thing').innerText;

            // Initialize the Showdown converter
            const converter = new showdown.Converter();

            // Convert Markdown to HTML
            const result = converter.makeHtml(input);

            // Set the HTML content of the result div
            document.getElementById('markdow').innerHTML = result;

            // Optionally, you can call other functions to process the output
            const markdown = document.getElementById('markdow');
            rec(markdown);
        }
//function renderMarkdown() {
//    const input =document.getElementById('thing').innerText;
//    const result = marked.parse(input);
//    document.getElementById('markdow').innerHTML = result;
//    const markdown = document.getElementById('markdow') ;
//    rec(markdown) ;
//
//}

renderMarkdown();
    
hljs.highlightAll() 
