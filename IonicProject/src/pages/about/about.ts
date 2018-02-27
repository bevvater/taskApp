import { Component } from '@angular/core';
import { NavController } from 'ionic-angular';

import { Storage } from '@ionic/storage';
import { LoginPage } from '../login/login';

@Component({
  selector: 'page-about',
  templateUrl: 'about.html'
})
export class AboutPage {


    private tokenStorageName = 'token';
    private userInfoStorageName = 'userInfo';

	constructor(public navCtrl: NavController,
		private storage: Storage) {

	}


  logout() {
  	// this.storage.remove(this.tokenStorageName);
  	// this.storage.remove(this.userInfoStorageName);
  	this.storage.clear();
 	// this.navCtrl.setRoot(LoginPage)
  	// this.navCtrl.push(LoginPage);
 	location.reload();
  }

}
