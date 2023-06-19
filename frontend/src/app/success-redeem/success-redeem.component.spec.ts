import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SuccessRedeemComponent } from './success-redeem.component';

describe('SuccessRedeemComponent', () => {
  let component: SuccessRedeemComponent;
  let fixture: ComponentFixture<SuccessRedeemComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [SuccessRedeemComponent]
    });
    fixture = TestBed.createComponent(SuccessRedeemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
