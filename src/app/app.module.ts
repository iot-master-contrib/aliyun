import { BaseModule } from './base/base.module';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NZ_I18N } from 'ng-zorro-antd/i18n';
import { zh_CN } from 'ng-zorro-antd/i18n';
import { APP_BASE_HREF, registerLocaleData } from '@angular/common';
import zh from '@angular/common/locales/zh'; 
import { CdkDrag, CdkDropList } from "@angular/cdk/drag-drop";
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AliyunComponent } from './setting/aliyun/aliyun.component';
import { TcCloudComponent } from './setting/tc-cloud/tc-cloud.component'; 
import { HicloudComponent } from './setting/hicloud/hicloud.component';
import { MsgLogComponent } from './msg-log/msg-log.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms'; 
import { NzLayoutModule } from 'ng-zorro-antd/layout';
import { NzMenuModule } from 'ng-zorro-antd/menu';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzCardModule } from 'ng-zorro-antd/card';
import { NzFormModule } from 'ng-zorro-antd/form'; 
import { NzInputModule } from 'ng-zorro-antd/input';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzInputNumberModule } from 'ng-zorro-antd/input-number';
import { NzSelectModule } from 'ng-zorro-antd/select';
import { NzUploadModule } from 'ng-zorro-antd/upload';
import { NzSwitchModule } from 'ng-zorro-antd/switch';
import { NzMessageModule } from 'ng-zorro-antd/message';
import { NzNotificationModule } from 'ng-zorro-antd/notification';
import { NzTableModule } from 'ng-zorro-antd/table';
import { NzTagModule } from 'ng-zorro-antd/tag';
import { NzModalModule } from 'ng-zorro-antd/modal';
import { NzPopconfirmModule } from 'ng-zorro-antd/popconfirm';
registerLocaleData(zh);

@NgModule({
  declarations: [
    AppComponent,
    AliyunComponent,
    TcCloudComponent,
    HicloudComponent,
    MsgLogComponent,
  ],
  imports: [
    BrowserModule,
    BaseModule,
    NzLayoutModule,
    NzMenuModule,
    NzPopconfirmModule,
    NzModalModule,
    NzTagModule,
    NzTableModule,
    NzIconModule,NzInputNumberModule,
    NzUploadModule,
    NzButtonModule,
    NzFormModule,
    NzInputModule ,
    NzSwitchModule,
    CdkDrag, CdkDropList,
    NzMessageModule,
    ReactiveFormsModule,
    NzNotificationModule,
    NzCardModule,
    NzSelectModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
  ],
  providers: [
    { provide: APP_BASE_HREF, useValue: '/app/sms/' },
    { provide: NZ_I18N, useValue: zh_CN }],
      bootstrap: [AppComponent],
})
export class AppModule {}
