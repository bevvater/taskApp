import { NavController, LoadingController, NavParams } from 'ionic-angular';
import { Component } from '@angular/core';

import { Observable } from 'rxjs/Rx';
import { LoginService } from './services/login.service';
import { User } from './user';
import { Storage } from '@ionic/storage';
import { RegisterPage } from '../register/register';
import { HomePage } from '../home/home';

import { TabsPage } from '../tabs/tabs';

@Component({
  selector: 'page-login',
  templateUrl: 'login.html'
})

export class LoginPage {


//    @Output() childEvent: EventEmitter<string> = new EventEmitter();
    
    constructor(
        private LoginService: LoginService,
        private storage: Storage,
		public navCtrl: NavController,
        public loadingCtrl: LoadingController    ) { }

    public model = new User('', '');
    private tokenStorageName = 'token';
    private userInfoStorageName = 'userInfo';

    submitLogin() {

        let loading = this.loadingCtrl.create({
            content: 'Please wait...'
        });
        loading.present();
        let userOperation: Observable<User[]>;
        userOperation = this.LoginService.getToken(this.model);
        userOperation.subscribe(
            obj => {

                this.storage.ready().then(() => {

                    console.log("lozzzzz 1")


                    this.storage.set(this.tokenStorageName, obj["token"]);
                    this.storage.set(this.userInfoStorageName, obj["user"]);

                    console.log(obj["token"]);
                    console.log(obj["user"]);

                    setTimeout(() => {
                        console.log("set time out!")
                    }, 5000);


                }).then(() => {
                    this.navCtrl.setRoot(TabsPage);
                    loading.dismiss();

                    console.log("lozzzzz 2")

                });


                // this.navCtrl.setRoot(TabsPage);


                // setTimeout(() => {
                //     this.navCtrl.setRoot(TabsPage);
                // }, 5000);



                // this.navCtrl.push(HomePage);
                // this.navCtrl.setRoot(TabsPage)

            },
            err => {
                console.log(err);
        });
    }

    register() {
        this.navCtrl.push(RegisterPage);
    }

}
