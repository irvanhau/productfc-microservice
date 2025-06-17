package service

import (
	"context"
	"productfc/cmd/product/repository"
	"productfc/models"
)

type ProductService struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(productService repository.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: productService,
	}
}

// Layer Service menentukan memakai DB atau Redis
func (s *ProductService) GetProductByID(ctx context.Context, productID int64) (*models.Product, error) {
	product, err := s.ProductRepository.FindProductByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) GetProductCategoryByID(ctx context.Context, productCategoryID int) (*models.ProductCategory, error) {
	productCategory, err := s.ProductRepository.FindProductCategoryByID(ctx, productCategoryID)
	if err != nil {
		return nil, err
	}

	return productCategory, nil
}

func (s *ProductService) CreateNewProduct(ctx context.Context, param *models.Product) (int64, error) {
	productID, err := s.ProductRepository.InsertNewProduct(ctx, param)
	if err != nil {
		return 0, err
	}

	return productID, nil
}

func (s *ProductService) CreateNewProductCategory(ctx context.Context, param *models.ProductCategory) (int, error) {
	productCategoryID, err := s.ProductRepository.InsertNewProductCategory(ctx, param)
	if err != nil {
		return 0, err
	}
	return productCategoryID, nil
}

func (s *ProductService) EditProduct(ctx context.Context, param *models.Product) (*models.Product, error) {
	product, err := s.ProductRepository.UpdateProduct(ctx, param)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) EditProductCategory(ctx context.Context, param *models.ProductCategory) (*models.ProductCategory, error) {
	productCategory, err := s.ProductRepository.UpdateProductCategory(ctx, param)
	if err != nil {
		return nil, err
	}

	return productCategory, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, productID int64) error {
	err := s.ProductRepository.DeleteProduct(ctx, productID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProductService) DeleteProductCategory(ctx context.Context, productCategoryID int) error {
	err := s.ProductRepository.DeleteProductCategory(ctx, productCategoryID)
	if err != nil {
		return err
	}

	return nil
}
