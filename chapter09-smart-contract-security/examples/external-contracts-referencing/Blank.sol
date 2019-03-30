pragma solidity ^0.5.6;

 contract Blank {
     event Print(string text);
     function () {
         emit Print("Here");
         // put malicious code here and it will run
     }
 }