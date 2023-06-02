import {SmsComponent} from './sms/sms.component';
import {AliyunComponent} from './setting/aliyun/aliyun.component';
import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {SubscribersComponent} from "./subscribers/subscribers.component";
import {SubscriberEditComponent} from "./subscriber-edit/subscriber-edit.component";

const routes: Routes = [
    {path: 'sms', component: SmsComponent},
    {path: 'subscriber', component: SubscribersComponent},
    {path: 'subscriber/edit/:id', component: SubscriberEditComponent},
    {path: 'subscriber/create', component: SubscriberEditComponent},
    {path: 'sms', component: SmsComponent},
    {path: 'setting', component: AliyunComponent},

];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule {
}
