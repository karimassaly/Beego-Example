function LoginRequest() {
    email = document.getElementById("Email").value;
    password = document.getElementById("password").value;
    
    $.ajax({
        url: 'http://localhost:8080/login',
        type: 'POST',
        data : 'Email=' + email + '&password=' + password,
        async: false,
        success: function () {
          window.location.href = "http://localhost:8080/main/";
        },
        error: function() {
          swal({title: "Failed", text: "Failed to Login", icon: 
          "error"}).then(function(){ 
              window.location.href = "http://localhost:8080/#signin";
                      }
          );
        }
    })
}