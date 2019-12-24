  $("a").on('click', function(event) {
      // Make sure this.hash has a value before overriding default behavior
      if (this.hash !== "") {
        // Prevent default anchor click behavior
        event.preventDefault();
  
        // Store hash
        var hash = this.hash;
  
        // Using jQuery's animate() method to add smooth page scroll
        // The optional number (800) specifies the number of milliseconds it takes to scroll to the specified area
        $('html, body').animate({
          scrollbot: $(hash).offset().bot
        }, 800, function(){
      
        // Add hash (#) to URL when done scrolling (default click behavior)
          window.location.hash = hash;
        });
      } // End if
    });


function check_both_pass() {
  re_pass = document.getElementById("re_pass").value
  password = document.getElementById("pass").value;

  console.log(re_pass, password)
  if (password != re_pass) {
    console.log("Im here now")
    swal({title: "Failed", text: "Password doesn't match the repeat password", icon: 
        "error"})
    return 
  } else {
    RegisterRequest()
  }
}

function RegisterRequest() {
      email = document.getElementById("email").value;
      password = document.getElementById("pass").value;
      name = document.getElementById("name").value;
      re_pass = document.getElementById("re_pass").value
      
      $.ajax({
          url: 'http://localhost:8080/register',
          type: 'POST',
          data : 'Email=' + email + '&password=' + password + "&name=" + name,
          async: false,
          success: function () {
            swal({title: "Good job!", text: "Account Successfully registered", icon: 
            "success"}).then(function(){ 
                location.reload();
                        }
            );
            window.location.href = "http://localhost:8080/#signin";
          },
          error: function() {
            swal({title: "Failed", text: "Account Already Exists", icon: 
            "error"}).then(function(){ 
                window.location.href = "http://localhost:8080/#";
                        }
            );
          }
      })
}