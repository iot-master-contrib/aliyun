import { MsgLogComponent } from './msg-log/msg-log.component';
import { TcCloudComponent } from './setting/tc-cloud/tc-cloud.component';
import { HicloudComponent } from './setting/hicloud/hicloud.component';
import { AliyunComponent } from './setting/aliyun/aliyun.component';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router'; 
 
const routes: Routes = [
 // { path: '', pathMatch: "full", redirectTo: "setting/aliyun" },
  { path: 'setting/aliyun', component: AliyunComponent },
  { path: 'setting/hicloud', component: HicloudComponent},
  { path: 'setting/tcCloud', component: TcCloudComponent},
  { path: 'msgLog', component: MsgLogComponent},
   
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
