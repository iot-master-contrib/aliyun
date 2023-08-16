import {SmsComponent} from './sms/sms.component';
import {AliyunComponent} from './setting/aliyun/aliyun.component';
import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';

const routes: Routes = [
    {path: 'sms', component: SmsComponent},
    {path: 'setting', component: AliyunComponent},

];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule {
}
