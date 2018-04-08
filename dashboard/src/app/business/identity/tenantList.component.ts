import { Component, OnInit, ViewContainerRef, ViewChild, Directive, ElementRef, HostBinding, HostListener } from '@angular/core';
import { I18NService } from 'app/shared/api';
import { AppService } from 'app/app.service';
import { I18nPluralPipe } from '@angular/common';
import { trigger, state, style, transition, animate } from '@angular/animations';
import { MenuItem } from '../../components/common/api';

@Component({
    selector: 'tenant-list',
    templateUrl: 'tenantList.html',
    animations: []
})
export class TenantListComponent implements OnInit {
    tenants = [];
    isDetailFinished = false;

    constructor(
        // private I18N: I18NService,
        // private router: Router
    ) { }

    ngOnInit() {
        this.tenants = [
            { "name": "tenant_A", "description": "--", },
            { "name": "tenant_B", "description": "--", }
        ];
    }

    onRowExpand(evt){
        this.isDetailFinished = false;
        console.log(evt.data.name);
        
    }

}
