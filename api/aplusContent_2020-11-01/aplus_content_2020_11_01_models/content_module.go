// Code generated by go-swagger; DO NOT EDIT.

package aplus_content_2020_11_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContentModule An A+ Content module. An A+ Content document is composed of content modules. The contentModuleType property selects which content module types to use.
//
// swagger:model ContentModule
type ContentModule struct {

	// content module type
	// Required: true
	ContentModuleType *ContentModuleType `json:"contentModuleType"`

	// standard company logo
	StandardCompanyLogo *StandardCompanyLogoModule `json:"standardCompanyLogo,omitempty"`

	// standard comparison table
	StandardComparisonTable *StandardComparisonTableModule `json:"standardComparisonTable,omitempty"`

	// standard four image text
	StandardFourImageText *StandardFourImageTextModule `json:"standardFourImageText,omitempty"`

	// standard four image text quadrant
	StandardFourImageTextQuadrant *StandardFourImageTextQuadrantModule `json:"standardFourImageTextQuadrant,omitempty"`

	// standard header image text
	StandardHeaderImageText *StandardHeaderImageTextModule `json:"standardHeaderImageText,omitempty"`

	// standard image sidebar
	StandardImageSidebar *StandardImageSidebarModule `json:"standardImageSidebar,omitempty"`

	// standard image text overlay
	StandardImageTextOverlay *StandardImageTextOverlayModule `json:"standardImageTextOverlay,omitempty"`

	// standard multiple image text
	StandardMultipleImageText *StandardMultipleImageTextModule `json:"standardMultipleImageText,omitempty"`

	// standard product description
	StandardProductDescription *StandardProductDescriptionModule `json:"standardProductDescription,omitempty"`

	// standard single image highlights
	StandardSingleImageHighlights *StandardSingleImageHighlightsModule `json:"standardSingleImageHighlights,omitempty"`

	// standard single image specs detail
	StandardSingleImageSpecsDetail *StandardSingleImageSpecsDetailModule `json:"standardSingleImageSpecsDetail,omitempty"`

	// standard single side image
	StandardSingleSideImage *StandardSingleSideImageModule `json:"standardSingleSideImage,omitempty"`

	// standard tech specs
	StandardTechSpecs *StandardTechSpecsModule `json:"standardTechSpecs,omitempty"`

	// standard text
	StandardText *StandardTextModule `json:"standardText,omitempty"`

	// standard three image text
	StandardThreeImageText *StandardThreeImageTextModule `json:"standardThreeImageText,omitempty"`
}

// Validate validates this content module
func (m *ContentModule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentModuleType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandardCompanyLogo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandardComparisonTable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandardFourImageText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandardFourImageTextQuadrant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandardHeaderImageText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandardImageSidebar(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandardImageTextOverlay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandardMultipleImageText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandardProductDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandardSingleImageHighlights(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandardSingleImageSpecsDetail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandardSingleSideImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandardTechSpecs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandardText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandardThreeImageText(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentModule) validateContentModuleType(formats strfmt.Registry) error {

	if err := validate.Required("contentModuleType", "body", m.ContentModuleType); err != nil {
		return err
	}

	if err := validate.Required("contentModuleType", "body", m.ContentModuleType); err != nil {
		return err
	}

	if m.ContentModuleType != nil {
		if err := m.ContentModuleType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contentModuleType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contentModuleType")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) validateStandardCompanyLogo(formats strfmt.Registry) error {
	if swag.IsZero(m.StandardCompanyLogo) { // not required
		return nil
	}

	if m.StandardCompanyLogo != nil {
		if err := m.StandardCompanyLogo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardCompanyLogo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardCompanyLogo")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) validateStandardComparisonTable(formats strfmt.Registry) error {
	if swag.IsZero(m.StandardComparisonTable) { // not required
		return nil
	}

	if m.StandardComparisonTable != nil {
		if err := m.StandardComparisonTable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardComparisonTable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardComparisonTable")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) validateStandardFourImageText(formats strfmt.Registry) error {
	if swag.IsZero(m.StandardFourImageText) { // not required
		return nil
	}

	if m.StandardFourImageText != nil {
		if err := m.StandardFourImageText.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardFourImageText")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardFourImageText")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) validateStandardFourImageTextQuadrant(formats strfmt.Registry) error {
	if swag.IsZero(m.StandardFourImageTextQuadrant) { // not required
		return nil
	}

	if m.StandardFourImageTextQuadrant != nil {
		if err := m.StandardFourImageTextQuadrant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardFourImageTextQuadrant")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardFourImageTextQuadrant")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) validateStandardHeaderImageText(formats strfmt.Registry) error {
	if swag.IsZero(m.StandardHeaderImageText) { // not required
		return nil
	}

	if m.StandardHeaderImageText != nil {
		if err := m.StandardHeaderImageText.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardHeaderImageText")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardHeaderImageText")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) validateStandardImageSidebar(formats strfmt.Registry) error {
	if swag.IsZero(m.StandardImageSidebar) { // not required
		return nil
	}

	if m.StandardImageSidebar != nil {
		if err := m.StandardImageSidebar.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardImageSidebar")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardImageSidebar")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) validateStandardImageTextOverlay(formats strfmt.Registry) error {
	if swag.IsZero(m.StandardImageTextOverlay) { // not required
		return nil
	}

	if m.StandardImageTextOverlay != nil {
		if err := m.StandardImageTextOverlay.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardImageTextOverlay")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardImageTextOverlay")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) validateStandardMultipleImageText(formats strfmt.Registry) error {
	if swag.IsZero(m.StandardMultipleImageText) { // not required
		return nil
	}

	if m.StandardMultipleImageText != nil {
		if err := m.StandardMultipleImageText.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardMultipleImageText")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardMultipleImageText")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) validateStandardProductDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.StandardProductDescription) { // not required
		return nil
	}

	if m.StandardProductDescription != nil {
		if err := m.StandardProductDescription.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardProductDescription")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardProductDescription")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) validateStandardSingleImageHighlights(formats strfmt.Registry) error {
	if swag.IsZero(m.StandardSingleImageHighlights) { // not required
		return nil
	}

	if m.StandardSingleImageHighlights != nil {
		if err := m.StandardSingleImageHighlights.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardSingleImageHighlights")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardSingleImageHighlights")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) validateStandardSingleImageSpecsDetail(formats strfmt.Registry) error {
	if swag.IsZero(m.StandardSingleImageSpecsDetail) { // not required
		return nil
	}

	if m.StandardSingleImageSpecsDetail != nil {
		if err := m.StandardSingleImageSpecsDetail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardSingleImageSpecsDetail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardSingleImageSpecsDetail")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) validateStandardSingleSideImage(formats strfmt.Registry) error {
	if swag.IsZero(m.StandardSingleSideImage) { // not required
		return nil
	}

	if m.StandardSingleSideImage != nil {
		if err := m.StandardSingleSideImage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardSingleSideImage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardSingleSideImage")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) validateStandardTechSpecs(formats strfmt.Registry) error {
	if swag.IsZero(m.StandardTechSpecs) { // not required
		return nil
	}

	if m.StandardTechSpecs != nil {
		if err := m.StandardTechSpecs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardTechSpecs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardTechSpecs")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) validateStandardText(formats strfmt.Registry) error {
	if swag.IsZero(m.StandardText) { // not required
		return nil
	}

	if m.StandardText != nil {
		if err := m.StandardText.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardText")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardText")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) validateStandardThreeImageText(formats strfmt.Registry) error {
	if swag.IsZero(m.StandardThreeImageText) { // not required
		return nil
	}

	if m.StandardThreeImageText != nil {
		if err := m.StandardThreeImageText.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardThreeImageText")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardThreeImageText")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this content module based on the context it is used
func (m *ContentModule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContentModuleType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandardCompanyLogo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandardComparisonTable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandardFourImageText(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandardFourImageTextQuadrant(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandardHeaderImageText(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandardImageSidebar(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandardImageTextOverlay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandardMultipleImageText(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandardProductDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandardSingleImageHighlights(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandardSingleImageSpecsDetail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandardSingleSideImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandardTechSpecs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandardText(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandardThreeImageText(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentModule) contextValidateContentModuleType(ctx context.Context, formats strfmt.Registry) error {

	if m.ContentModuleType != nil {
		if err := m.ContentModuleType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contentModuleType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contentModuleType")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) contextValidateStandardCompanyLogo(ctx context.Context, formats strfmt.Registry) error {

	if m.StandardCompanyLogo != nil {
		if err := m.StandardCompanyLogo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardCompanyLogo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardCompanyLogo")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) contextValidateStandardComparisonTable(ctx context.Context, formats strfmt.Registry) error {

	if m.StandardComparisonTable != nil {
		if err := m.StandardComparisonTable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardComparisonTable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardComparisonTable")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) contextValidateStandardFourImageText(ctx context.Context, formats strfmt.Registry) error {

	if m.StandardFourImageText != nil {
		if err := m.StandardFourImageText.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardFourImageText")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardFourImageText")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) contextValidateStandardFourImageTextQuadrant(ctx context.Context, formats strfmt.Registry) error {

	if m.StandardFourImageTextQuadrant != nil {
		if err := m.StandardFourImageTextQuadrant.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardFourImageTextQuadrant")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardFourImageTextQuadrant")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) contextValidateStandardHeaderImageText(ctx context.Context, formats strfmt.Registry) error {

	if m.StandardHeaderImageText != nil {
		if err := m.StandardHeaderImageText.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardHeaderImageText")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardHeaderImageText")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) contextValidateStandardImageSidebar(ctx context.Context, formats strfmt.Registry) error {

	if m.StandardImageSidebar != nil {
		if err := m.StandardImageSidebar.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardImageSidebar")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardImageSidebar")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) contextValidateStandardImageTextOverlay(ctx context.Context, formats strfmt.Registry) error {

	if m.StandardImageTextOverlay != nil {
		if err := m.StandardImageTextOverlay.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardImageTextOverlay")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardImageTextOverlay")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) contextValidateStandardMultipleImageText(ctx context.Context, formats strfmt.Registry) error {

	if m.StandardMultipleImageText != nil {
		if err := m.StandardMultipleImageText.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardMultipleImageText")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardMultipleImageText")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) contextValidateStandardProductDescription(ctx context.Context, formats strfmt.Registry) error {

	if m.StandardProductDescription != nil {
		if err := m.StandardProductDescription.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardProductDescription")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardProductDescription")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) contextValidateStandardSingleImageHighlights(ctx context.Context, formats strfmt.Registry) error {

	if m.StandardSingleImageHighlights != nil {
		if err := m.StandardSingleImageHighlights.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardSingleImageHighlights")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardSingleImageHighlights")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) contextValidateStandardSingleImageSpecsDetail(ctx context.Context, formats strfmt.Registry) error {

	if m.StandardSingleImageSpecsDetail != nil {
		if err := m.StandardSingleImageSpecsDetail.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardSingleImageSpecsDetail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardSingleImageSpecsDetail")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) contextValidateStandardSingleSideImage(ctx context.Context, formats strfmt.Registry) error {

	if m.StandardSingleSideImage != nil {
		if err := m.StandardSingleSideImage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardSingleSideImage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardSingleSideImage")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) contextValidateStandardTechSpecs(ctx context.Context, formats strfmt.Registry) error {

	if m.StandardTechSpecs != nil {
		if err := m.StandardTechSpecs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardTechSpecs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardTechSpecs")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) contextValidateStandardText(ctx context.Context, formats strfmt.Registry) error {

	if m.StandardText != nil {
		if err := m.StandardText.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardText")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardText")
			}
			return err
		}
	}

	return nil
}

func (m *ContentModule) contextValidateStandardThreeImageText(ctx context.Context, formats strfmt.Registry) error {

	if m.StandardThreeImageText != nil {
		if err := m.StandardThreeImageText.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standardThreeImageText")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standardThreeImageText")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentModule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentModule) UnmarshalBinary(b []byte) error {
	var res ContentModule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}