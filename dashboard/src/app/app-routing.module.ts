import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

const routes: Routes = [
    {path: '', redirectTo: '/home', pathMatch: 'full'},
    {path: 'home', loadChildren: './business/home/home.module#HomeModule'},
    {path: 'service', loadChildren: './business/service/service.module#ServiceModule'},
    {path: 'block', loadChildren: './business/block/block.module#BlockModule'},
    {path: 'createVolume', loadChildren: './business/block/create-volume/create-volume.module#CreateVolumeModule'},
    {path: 'cloud', loadChildren: './business/cloud/cloud.module#CloudModule'},
    {path: 'profile', loadChildren: './business/profile/profile.module#ProfileModule'},
    {path: 'createProfile', loadChildren: './business/profile/createProfile/createProfile.module#CreateProfileModule'},
    {path: 'modifyProfile/:profileId', loadChildren: './business/profile/modifyProfile/modifyProfile.module#ModifyProfileModule'},
    {path: 'resource', loadChildren: './business/resource/resource.module#ResourceModule'},
    {path: 'identity', loadChildren: './business/identity/identity.module#IdentityModule'}
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule {}
