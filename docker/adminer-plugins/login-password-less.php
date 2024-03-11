<?php

require_once('plugins/login-password-less.php');
/** Set allowed password
 * @param string result of password_hash
 */
return new AdminerLoginPasswordLess(
  $password_hash = password_hash("admin", PASSWORD_DEFAULT)
);


// require_once('plugins/login-password-less.php');
// /** Set allowed password
// 	* @param string result of password_hash
// 	*/
// return new AdminerLoginPasswordLess(
// 	$password_hash = password_hash("admin", "543216")
// );

// require_once('plugins/login-password-less.php');
// /** Set allowed password
//  * @param string result of password_hash
//  */
// return new AdminerLoginPasswordLess(
//     // mkpasswd -m SHA-512 $PASSWORD
//     // $password_hash = ""
//     $password_hash = "$6$kK8A7sjb5vbFNURy$/uPoYblVzJ0FFhYvfajdhbC/eRg.x9/k17rCoDpG0jdSSOdkexgTTx.FIE938biqPZvyCFYGmZcFMEeLhy7Bc/"
//     // function login($login, $password) {
//     //         return true;
//     // }
// );

// require_once('plugins/login-password-less.php');
// /** Set allowed password
//  * @param string result of password_hash
//  */
// class MyAdminerLoginPasswordLess extends AdminerLoginPasswordLess {
//     function login($login, $password) {
//         // Here, always return true to bypass password authentication
//         // In a real scenario, you might want to add some logic to check
//         // whether $login is in a list of allowed usernames, etc.
//         return true;
//     }
// }
// // Create an instance of your custom AdminerLoginPasswordLess
// return new MyAdminerLoginPasswordLess();
