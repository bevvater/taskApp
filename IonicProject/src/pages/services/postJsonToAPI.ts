// /* * * ./app/comments/services/comment.service.ts * * */
// import { Injectable } from '@angular/core';
// import { Http, Response, Headers, RequestOptions } from '@angular/http';
// import { User } from '../user';
// import {Observable} from 'rxjs/Rx';
// import 'rxjs/add/operator/map';
// import 'rxjs/add/operator/catch';

// @Injectable()
// export class PostJsonToApiService {
// 	constructor (private http: Http) {}

// 	getToken(body: Object): Observable<User[]> {
//     const bodyString = JSON.stringify(body);
    
//     // console.log(bodyString)
//     // console.log(typeof bodyString)

// 	const headers = new Headers({ 'Content-Type': 'application/json'});
// //	const headers = new Headers({ 'Content-Type': 'x-www-form-urlencoded'});
// //	const headers = new Headers({ 'Content-Type': 'application/form-data'});
//     const options = new RequestOptions({ headers: headers });
// 	return this.http.post('http://localhost:8080/login', bodyString, options)
// 		.map((res: Response) => res.json())
// 		.catch((error: any) => Observable.throw(error.json().error || 'Server error'));
// 	}

// }
