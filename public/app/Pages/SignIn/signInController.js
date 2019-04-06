'use strict';

var SignIn = angular.module('SignIn', ['ui.router'])

SignIn.config(['$routeProvider', function ($routeProvider) {
    $routeProvider.when('/SignIn', {
        templateUrl: 'Pages/SignIn/signIn.html',
        controller: 'SignInController',
        controllerAs: '$ctrl',
    });
}])

/**
 * @brief SignIn controller 
 */
SignIn.controller('SignInController', function ($location, $http) {

    this.newUser = {
        login: "",
        password: "",
        confirmPassword: "",
        name: "",
        surname: "",
        birthDate: "",
        country: "",
        city: "",
    }
    document.getElementById('test').style.display = "none";

    /**
     * @brief Cancel button handler
     */
    this.cancel = () => {
        console.log("cancel account")
        $location.path('/');
    }

    /**
     * @brief validation button handler
     */
    this.validate = () => {
        console.log("validation account")
        console.log(this.newUser)

        console.log(this.image)
        $http.post('/sign-in', {
            login: this.newUser.login,
            password: this.newUser.password,
            name: this.newUser.name,
            surname: this.newUser.surname,
            birthDate: this.newUser.birthDate,
            country: this.newUser.country,
            city: this.newUser.city,
            image: this.image,
        }).then(response => response.data)
        .then((response) => {
            console.log(response)
            $location.path('/Connection');
        }).catch((err) => console.error(err))
    }

    this.test = undefined;
    this.tmpFunction = () => {
        console.log(this.test)
    }

    this.image = undefined;
    this.handleFileSelect = (evt) => {
        var files = evt.target.files; // FileList object
        for (var i = 0, f; f = files[i]; i++) {
          if (!f.type.match('image.*')) {
            continue;
          }
          var reader = new FileReader();
          reader.onload = ((theFile) => {
            return (e) => {
              var span = document.createElement('span');
              console.log(e.target.result)
              this.image = e.target.result;
            };
          })(f);
          reader.readAsDataURL(f);
        }
      }
      document.getElementById('files').addEventListener('change', this.handleFileSelect, false);
});