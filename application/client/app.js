'use strict';

var app = angular.module('identityManagementApp', []);

app.controller('MainController', ['$scope', '$http', function($scope, $http) {
    // 새로운 신원 정보
    $scope.newIdentity = {};

    // 신원 정보 등록 함수
    $scope.registerIdentity = function() {
        $http.post('http://localhost:3000/CreateIdentity', $scope.newIdentity)
            .then(function(response) {
                alert('신원 정보가 등록되었습니다.');
                $scope.newIdentity = {}; // 입력 폼 초기화
            })
            .catch(function(error) {
                console.error('Failed to register identity:', error);
                alert('신원 정보 등록에 실패했습니다.');
            });
    };

    // 신원 정보 조회 함수
    $scope.queryIdentity = function() {
        $http.get('/QueryIdentity?idnumber=p1' )
            .then(function(response) {
                // 조회된 신원 정보를 사용하여 UI 업데이트 등의 작업 수행
                console.log('Queried identity:', response.data);
            })
            .catch(function(error) {
                console.error('Failed to query identity:', error);
                alert('신원 정보 조회에 실패했습니다.');
            });
    };
}]);
