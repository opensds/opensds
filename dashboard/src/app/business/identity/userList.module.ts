import { NgModule, APP_INITIALIZER } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';
import { UserListComponent } from './userList.component';
import { FormModule, ButtonModule, CheckboxModule, DataTableModule, DropMenuModule, DialogModule, InputTextModule, InputTextareaModule, DropdownModule, PasswordModule, ConfirmDialogModule } from '../../components/common/api';

import { UserDetailModule } from './userDetail/userDetail.module';

@NgModule({
  declarations: [UserListComponent],
  imports: [
    CommonModule,
    ButtonModule,
    DataTableModule,
    DropMenuModule,
    DialogModule,
    ConfirmDialogModule,
    InputTextModule,
    InputTextareaModule,  
    DropdownModule,
    CheckboxModule,
    PasswordModule,
    FormsModule,
    ReactiveFormsModule,
    FormModule,
    UserDetailModule
  ],
  exports: [UserListComponent],
  providers: []
})
export class UserListModule { }