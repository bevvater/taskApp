import { Component } from '@angular/core';
// import { NavController, LoadingController, NavParams } from 'ionic-angular';
import { NavController, LoadingController } from 'ionic-angular';

import { Storage } from '@ionic/storage';

@Component({
  selector: 'page-home',
  templateUrl: 'home.html'
})
export class HomePage {

	private fullname = "";
	private role = "";

	private lol = "......";

  constructor(public navCtrl: NavController, private storage: Storage,
  	public loadingCtrl: LoadingController) {


  		//this.lol = this.navParams.get('id');
		// let loading = this.loadingCtrl.create({
		// 	content: 'Please wait...'
		// });

		// loading.present();

		// setTimeout(() => {
		// 	loading.dismiss();
		// }, 5000);

  }

  ngOnInit() {
  	console.log("lozzzzz 3");

    this.storage.ready().then(() => {
    	console.log("lozzzzz 4");
        setTimeout(() => {
			this.storage.get('userInfo').then((val) => {
				if(val != null)
				{
					console.log("not null");
					this.fullname = val["fullname"];
					this.role = val["role"];	
				}
			});
        }, 1000);
    });

	}
}
