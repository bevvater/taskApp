import { Component } from '@angular/core';
import { NavController, LoadingController } from 'ionic-angular';
import { Task } from './taskModel';

@Component({
  selector: 'page-task',
  templateUrl: 'task.html'
})
export class TaskPage {

	constructor(public navCtrl: NavController,
		public loadingCtrl: LoadingController) {

  }


  ngOnInit() {

	let loading = this.loadingCtrl.create({
		content: 'Please wait...'
	});
	loading.present();

	setTimeout(() => {
		loading.dismiss();
	}, 5000);


  }

}
