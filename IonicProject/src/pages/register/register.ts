import { NavController } from 'ionic-angular';
import { Component} from '@angular/core';

import { Observable } from 'rxjs/Rx';
import { RegisterService } from './services/register.service';
import { User } from './user';

import { LoginPage } from '../login/login';
// import { Storage } from '@ionic/storage';

@Component({
  selector: 'page-register',
  templateUrl: 'register.html'
})

export class RegisterPage {
    constructor(
        private RegisterService: RegisterService,
        // private storage: Storage,
		public navCtrl: NavController    ) { }

    public model = new User("", "", "");
    // private tokenStorageName = 'token';

    submitRegister() {
        let userOperation : Observable<User[]>;
        userOperation = this.RegisterService.registerUser(this.model);
        userOperation.subscribe(
            obj => {
                console.log(obj);
            });
        this.navCtrl.push(LoginPage);
    }


    // submitLogin() {
    //     let userOperation: Observable<User[]>;
    //     userOperation = this.LoginService.getToken(this.model);
    //     userOperation.subscribe(
    //         obj => {
    //             this.storage.set(this.tokenStorageName, obj["token"]);
    //         },
    //         err => {
    //             console.log(err);
    //     });
    // }

}
