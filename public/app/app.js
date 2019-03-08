'use strict';

// Declare app level module which depends on views, and core components
angular.module('myApp', [
  'ngRoute',
  'ui.router',
  'Dashboard',
  'Connection',
  'Configuration',
  'SignIn',
  'Profile',
  'Details',
]).
config(['$routeProvider', function($routeProvider) {
  $routeProvider.otherwise({redirectTo: '/'});
}]);
